package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	//每一秒打印一个通知，说还在休眠
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	//休眠2分钟
	reportNap("sending gorouting", 2)
	fmt.Println("***sending value***")
	//在main仍处于休眠状态时，阻塞此发送
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	//休眠5秒
	reportNap("receiving gorouting", 5)

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
