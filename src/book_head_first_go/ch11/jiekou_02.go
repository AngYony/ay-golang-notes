package main

import (
	"book_head_first_go/ch11/mypkg"
	"fmt"
)

func main() {
	//声明一个接口类型的变量
	var value mypkg.MyInterface
	//将MyType类型的值赋值给该接口变量（MyType满足MyInterface接口）
	value = mypkg.MyType(5)
	//调用满足该接口的对应的方法
	value.MethodWithoutParameters()
	value.MethodWithParameter(11.1)
	fmt.Println(value.MethodWithReturnValue())

}
