package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	//每一秒打印一个通知，说还在休眠
	for i := 0; i < delay; i++ {
		fmt.Println(name, "正在休眠")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "休眠结束")
}

func send(myChannel chan string) {
	//休眠2秒
	reportNap("发送前休眠2秒", 2)
	fmt.Println("***sending value a***")
	myChannel <- "a" //发送值，阻塞当前goroutine，直到其他goroutine接收该channel的值
	fmt.Println("***sending value b***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel) //将以异步方式运行
	//休眠5秒
	reportNap("5秒休眠", 5)
	//直到5秒之后，才接收值，从而解除send goroutine中的阻塞
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
