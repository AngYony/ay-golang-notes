package main

import "fmt"

func main() {
	defer func() {
		//从惊恐中恢复
		if e := recover(); e != nil {
			fmt.Println(e) //输出：你好吗？
		}
	}()

	panic("你好吗？") //引发惊恐
}
