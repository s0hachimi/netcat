package function

import (
	"log"
	"os"
)

var Birti9 = Read("nc.txt")

func Read(s string) []byte {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
