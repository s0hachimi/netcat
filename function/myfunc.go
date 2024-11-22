package function

import (
	"fmt"
	"log"
	"net"
	"time"
)

func SendMessage(message, nameClient string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clientFormat := ""

	for client, name := range Clients {
		clientFormat = fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), name)
		if client != sender {
			client.Write([]byte("\n"))
			Message2clents := fmt.Sprintf("[%v][%v]:%v", T.Format(time.DateTime), nameClient, (message))
			_, err := client.Write([]byte(Message2clents))
			if err != nil {
				log.Println(err)
				removeClient(client)
				return
			}
			client.Write([]byte(clientFormat))
		}
	}
}

func JoinOrLeft(name, msg string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clientFormat := ""

	for cl, myName := range Clients {
		clientFormat = fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), myName)
		if cl != sender {
			History += name + " has " + msg + " our chat... \n"
			cl.Write([]byte("\n" + name + " has " + msg + " our chat... \n"))
			cl.Write([]byte(clientFormat))
		}
	}
}

// Lee has left our chat...

func isNameTaken(name string) bool {
	mu.Lock()
	defer mu.Unlock()

	for _, eName := range Clients {
		if eName == name {
			return true
		}
	}
	return false
}

func removeClient(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(Clients, conn)
	conn.Close()
}
