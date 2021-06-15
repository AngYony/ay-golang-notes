package main

import (
	"fmt"
	"time"
)

func a() {
	//使用循环打印500个字母a
	for i := 0; i < 500; i++ {
		fmt.Print("a")
	}
}

func b() {
	//使用循环打印500个字母b
	for i := 0; i < 500; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	//暂停main goroutine 3秒
	time.Sleep(time.Second * 3)

	fmt.Println("结束")
}
