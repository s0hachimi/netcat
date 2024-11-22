package function

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	T             = time.Now()
	f             bool
	Clients       = make(map[net.Conn]string)
	clientsJoined = make(map[net.Conn]bool)
	mu            sync.Mutex
	History       string
)

func HandleChat(client net.Conn) {
	defer removeClient(client)

	if len(Clients) == 10 {
		fmt.Fprintln(client, "The chat is already full, you are not allowed to enter it.")
		return
	}

	client.Write(Birti9)

	scanner := bufio.NewScanner(client)

	f = true
	var name string

start:

	for {

		for f {

			_, err := client.Write([]byte("[ENTER YOUR NAME]:"))
			if err != nil {
				fmt.Println("Error reading from input:", err)
				return
			}

			scanner.Scan()
			name = scanner.Text()

			for _, r := range name {
				if r < 32 || r > 127 {
					fmt.Println(r)
					fmt.Fprintln(client, "Invalid Name")
					goto start
				}
			}

			if name == "" || len(name) > 50 {
				fmt.Fprintln(client, "Invalid Name")
				goto start
			}

			if !isNameTaken(name) {
				mu.Lock()
				Clients[client] = name
				mu.Unlock()
			} else if name != "" {
				client.Write([]byte("This name is already taken! Please enter a different name.\n"))
				continue
			}

			if name != "" {
				fmt.Fprint(client, History)
				f = false
			}

		}

		if !clientsJoined[client] {
			JoinOrLeft(name, "joined", client)
		}

		mu.Lock()
		clientsJoined[client] = true
		mu.Unlock()

		Message := fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), name)

		_, err := client.Write([]byte(Message))
		if err != nil {
			fmt.Println("ww", err)
			return
		}

		buf := make([]byte, 2048)

		n, err := client.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Client closed connection gracefully.", err)
				removeClient(client)
				JoinOrLeft(name, "left", client)
				return
			}
		}

		if string(buf[:n]) == "\n" {
			continue
		}

		if string(buf[:n]) == "--name\n" {
			fmt.Fprintln(client, "You can change your name")
			f = true
			goto start
		}

		// fmt.Print(buf[:n], []byte("--name\n"))

		Message += string(buf[:n])
		History += Message

		SendMessage(string(buf[:n]), name, client)

		// fmt.Println(clients)

	}
}
