package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 500; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 500; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	//暂停main goroutine 1秒
	time.Sleep(time.Second * 3)

	fmt.Println("结束")
}
