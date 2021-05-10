package main

import "book.go_gjbc/ch1/wyfile"

func main() {

	var CloseFile = (*wyfile.File).Close

	var ReadFile = (*wyfile.File).Read

	f, _ := wyfile.OpenFile("wy.txt")
	ReadFile(f, 0, nil)
	CloseFile(f)
}
