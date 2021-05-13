package main

import "fmt"

func say(names []string) []string {
	return append(names, "AA", "BB")
}
func main() {
	// var my []string
	// fmt.Println(my == nil) //输出：true

	// //将nil直接传递给形参为切片的函数
	// abc := say(nil)  //不会报错
	// fmt.Println(abc) //正常输出：[AA BB]

	// var no *int      //定义一个指针类型的指针变量，但没有赋值
	// fmt.Println(no)  //输出：<nil>
	// fmt.Println(*no) //报错

	// var fn func(a, b int) int
	// fmt.Println(fn == nil) //输出：true

	// var soup map[string]int
	// fmt.Println(soup == nil) //输出：true

	// //读取操作，不报错
	// wy, ok := soup["wy"]
	// if ok {
	// 	fmt.Println("存在wy")
	// } else {
	// 	fmt.Println(wy) //输出int类型的零值：0
	// 	fmt.Println("不存在，写入看看")
	// 	//写入操作，报错
	// 	soup["wy"] = 1
	// }

	var v interface{}
	fmt.Printf("%T %v %v \n", v, v, v == nil) //输出：<nil> <nil> true

	var p *int
	v = p
	fmt.Printf("%T %v %v", v, v, v == nil) //输出：*int <nil> false
	fmt.Printf("%#v", v)                   //输出：(*int)(nil)
}
