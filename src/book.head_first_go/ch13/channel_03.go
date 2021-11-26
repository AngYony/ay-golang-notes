package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "Hi" //通过channel发送一个值
}

func main() {
	//创建一个新的channel
	myChannel := make(chan string)
	//将channel传递给新goroutine中运行的函数
	go greeting(myChannel)
	//从channel中接收值
	chv := <-myChannel
	fmt.Println(chv)
}
