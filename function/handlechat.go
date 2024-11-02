package function

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func HandleChat(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte(Birti9))

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
