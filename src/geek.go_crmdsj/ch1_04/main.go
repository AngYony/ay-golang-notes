package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	// 获取命令行输入的参数
	if len(os.Args) > 1 {
		fmt.Println("hello world!", os.Args[1])
	}
}
