package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	Port := ":8989"
	if len(os.Args) == 2 {
		Port = ":" + os.Args[1]
	}

	ln, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port " + Port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleChat(conn)

	}
}

func Read(s string) {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(file))
}

func HandleChat(conn net.Conn) {

	
	var name string
	// var chat string
	var c int

	T := time.Now()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if c == 0 {
			name = scanner.Text()
		}

		// chat = scanner.Text()

		fmt.Printf("[%v][%v]:", T.Format(time.DateTime), name)
		c++

		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(status)

	}
}
