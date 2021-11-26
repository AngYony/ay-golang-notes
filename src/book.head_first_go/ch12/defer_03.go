package main

import "fmt"

func socialize() {
	fmt.Println("拜拜")
	defer fmt.Println("再见") //该调用被推迟到socialize退出之后
	fmt.Println("滚蛋")
}
func main() {
	socialize()
}
