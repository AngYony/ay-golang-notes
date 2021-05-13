package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...", id, " 睡眠中 ...")
	c <- id //发送值
}

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	fmt.Println("已完成goroutine的全部启动")

	//time.After函数返回一个通道
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c: //从通道中取值
			fmt.Println("gopher ", gopherId, "完成睡眠")
		case <-timeout:
			fmt.Println("等待超时")
		}
	}
}
