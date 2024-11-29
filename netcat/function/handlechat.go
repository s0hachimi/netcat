package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	Clients       = make(map[net.Conn]string)
	clientsJoined = make(map[net.Conn]bool)
	mu            sync.Mutex
	History       string
	buf           = make([]byte, 2048)
)

const Max = 2

func HandleChat(client net.Conn) {
	defer removeClient(client)



	

	client.Write(Birti9)

	scanner := bufio.NewScanner(client)

	f := false
	var name string

agian:

	for {

		for !f {

			client.Write([]byte("[ENTER YOUR NAME]:"))

			scanner.Scan()
			mu.Lock()
			name = scanner.Text()
			mu.Unlock()

			for _, r := range name {
				if r < 32 || r > 127 {
					fmt.Fprintln(client, "Invalid Name")
					goto agian
				}
			}

			if name == "" || len(name) > 50 {
				fmt.Fprintln(client, "Invalid Name")
				goto agian
			}

			if !isNameTaken(name) {
				mu.Lock()
				Clients[client] = name
				mu.Unlock()
			} else if name != "" {
				client.Write([]byte("This name is already taken! Please enter a different name.\n"))
				goto agian
			}

			if len(Clients) > Max {
				fmt.Fprintln(client, "The chat is already full, you are not allowed to enter it.")
				return
			}

			if name != "" {
				fmt.Fprint(client, History)
				f = true
			}

		}
		
		

		if !clientsJoined[client] {
			JoinOrLeft(name, "joined", client)
		}

		mu.Lock()
		clientsJoined[client] = true
		mu.Unlock()

		Message := fmt.Sprintf("[%s][%s]:", time.Now().Format(time.DateTime), name)

		client.Write([]byte(Message))

		n, err := client.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Client closed connection gracefully.", err)
				removeClient(client)
				JoinOrLeft(name, "left", client)
				return
			}
		}

		// fmt.Println(string(buf[:n]))

		skip := false

		for _, r := range string(buf[:n]) {
			if r == 27 || r == 12 {
				skip = true
				continue
			}
		}

		if skip {
			continue
		}

		if string(buf[:n]) == "\n" {
			continue
		}

		if string(buf[:n]) == "--name\n" {
			fmt.Fprintln(client, "You can change your name")
			f = false
			goto agian
		}

		// fmt.Print(buf[:n], []byte("--name\n"))

		Message += string(buf[:n])
		History += Message

		SendMessage(string(buf[:n]), name, client)

		// fmt.Println(clients)

	}
}
