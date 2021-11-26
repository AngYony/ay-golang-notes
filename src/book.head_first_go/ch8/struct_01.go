package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 1     //设置struct字段值
	myStruct.word = "hello" //获取struct字段值
	fmt.Println(myStruct.number, myStruct.word)
}
