package function

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func HandleChat(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte(Birti9))

	var name string
	// var m string
	var f bool

	T := time.Now()

	scanner := bufio.NewScanner(conn)

	

	for scanner.Scan() {
		if !f {
			name = scanner.Text()
		}
		f = true

		Message := fmt.Sprintf("[%v][%v]:", T.Format(time.DateTime), name)

		_, err := conn.Write([]byte(Message))
		if err != nil {
			log.Println("Error reading from input:", err)
		}

		// m = scanner.Text()

		// status, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }

	}
}
