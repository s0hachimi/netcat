package function

import (
	"log"
	"os"
)

var Birti9 = Read("nc.txt")

func Read(s string) string {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
