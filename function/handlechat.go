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
	clients       = make(map[net.Conn]string)
	clientsJoined = make(map[net.Conn]bool)
	mu            sync.Mutex
)

func HandleChat(client net.Conn) {
	defer removeClient(client)

	client.Write([]byte(Birti9))

	scanner := bufio.NewScanner(client)

	f = false
	var name string

	for {
		for !f {
		start:
			_, err := client.Write([]byte("[ENTER YOUR NAME]:"))
			if err != nil {
				fmt.Println("Error reading from input:", err)
				return
			}
			scanner.Scan()
			name = scanner.Text()

			if name == "" {
				goto start
			}

			if !isNameTaken(name) {

				mu.Lock()
				clients[client] = name
				mu.Unlock()

			} else if name != "" {
				client.Write([]byte("This name is already taken! Please enter a different name.\n"))
				continue
			}

			if name != "" {
				f = true
			}

		}

		if !clientsJoined[client] {
			Join(name, client)
		}

		mu.Lock()
		clientsJoined[client] = true
		mu.Unlock()

		Message := fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), name)

		_, err := client.Write([]byte(Message))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 2000)

		_, err = client.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Client closed connection gracefully.")
			}
		}

		SendMessage(string(buf), name, client)

	}
}

func SendMessage(message, nameClient string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clientFormat := ""

	for client, name := range clients {
		clientFormat = fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), name)
		if client != sender {
			client.Write([]byte("\n"))
			Message2clents := fmt.Sprintf("[%v][%v]:%v", T.Format(time.DateTime), nameClient, (message))
			_, err := client.Write([]byte(Message2clents))
			if err != nil {
				fmt.Println(err)
				delete(clients, client)
				client.Close()
			}
			client.Write([]byte(clientFormat))
		}
	}
}

func Join(name string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clientFormat := ""

	for cl, myName := range clients {
		clientFormat = fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), myName)
		if cl != sender {
			cl.Write([]byte("\n" + name + " has joined our chat... \n"))
			cl.Write([]byte(clientFormat))
		}
	}
}

func isNameTaken(name string) bool {
	mu.Lock()
	defer mu.Unlock()

	for _, eName := range clients {
		if eName == name {
			return true
		}
	}
	return false
}

func removeClient(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, conn)
}
