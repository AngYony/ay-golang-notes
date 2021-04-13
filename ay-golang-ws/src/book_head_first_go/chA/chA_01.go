package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_RDONLY
	file, err := os.OpenFile("wy.txt", options, os.FileMode(0600))
}
