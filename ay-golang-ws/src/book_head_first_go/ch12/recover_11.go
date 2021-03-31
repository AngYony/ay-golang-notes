package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func freakOut() {
	defer calmDown() //延迟对函数的调用
	panic("崩溃了。。。")
	fmt.Println("panic之后的代码永远不会被执行")
}

func main() {
	freakOut()
	//这段代码在freakOut2返回之后会运行
	fmt.Println("运行完成")
}
