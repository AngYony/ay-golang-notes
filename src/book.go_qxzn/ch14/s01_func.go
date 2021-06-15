package main

import "fmt"

//定义一个新的函数类型sayF
type sayF func() string

//形式是sayF类型的变量，并且返回的也是sayF类型函数
func hello(s sayF, riyu string) sayF {
	return func() string {
		return s() + riyu
	}
}

func english() string {
	return "英语"
}

func main() {
	//获取返回的函数赋值给变量
	lb := hello(english, "日语")
	//执行变量对应的函数体
	fmt.Println(lb())
}
