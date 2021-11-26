package main

import "fmt"

func syaHi() {
	fmt.Println("hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}
func main() {
	var myFunc func()
	myFunc = syaHi
	myFunc()

	twice(sayBye)
}
