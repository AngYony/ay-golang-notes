package main

import "fmt"

func greeting2(myChannel chan string) {
	myChannel <- "hi" //发送操作会导致该goroutine阻塞
}
func main() {
	myChannel := make(chan string)
	//在main goroutine中发送值，会阻塞main goroutine，从而引发异常
	myChannel <- "你看"
	//由于阻塞，不会被输出
	fmt.Println("被阻塞了")
	//go语句由于上述的阻塞，不会被执行
	go greeting2(myChannel)
	fmt.Println(<-myChannel)
}
