package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"netcat/function"
)

func main() {
	Port := ":8989" // 65535

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(os.Args) == 2 {
		Port = ":" + os.Args[1]
	}

	if Port == ":" || Port == ":0" {
		log.Fatal("This Port not allowed !!!")
	}

	ln, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on the port " + Port)

	count := 0

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		count++

		if count == 11 {
			conn.Write([]byte("you have no accept !!!"))
			return
		}

		
		go function.HandleChat(conn)

	}
}
