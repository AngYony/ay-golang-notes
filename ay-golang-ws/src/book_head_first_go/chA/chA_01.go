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
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("wy.txt", options, os.FileMode(0600))
}
