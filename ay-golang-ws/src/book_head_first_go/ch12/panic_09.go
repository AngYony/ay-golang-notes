package main

import "fmt"

func one() {
	//这个函数调用首先被延迟，所以它将在最后执行。
	defer fmt.Println("defer-one")
	two()
}

func two() {
	//这个函数调用最后被延迟，所以它将首先执行
	defer fmt.Println("defer-two")
	panic("two引发了崩溃")
}

func main() {
	one()
}
