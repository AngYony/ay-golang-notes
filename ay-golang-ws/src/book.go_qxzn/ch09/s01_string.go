package main

import "fmt"

func main() {

	// var wypath = `C:\go`
	// fmt.Print(wypath)

	A1 := 'A'
	var B rune = 'B'
	//获取单个字符的代码点
	fmt.Printf("%T %[1]d  \n", A1) //输出：int32 65
	fmt.Printf("%T \n", B)         //输出：int32

	//根据代码点创建变量
	var A2 rune = 65
	var pi rune = 960 //π的代码点
	//获取对应表示的字符
	fmt.Printf("%c %c", A2, pi) //输出：A π

}
