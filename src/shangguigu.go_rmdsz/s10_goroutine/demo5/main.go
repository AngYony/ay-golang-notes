package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func sayFail() {
	// defer必须定义在方法开头
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("方法发生了错误,", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "故意引发panic"
	// panic()

}
func main() {
	go sayHello()
	go sayFail()
	for {
		fmt.Println("测试", time.Now())
		time.Sleep(time.Second)
	}
}
