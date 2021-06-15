package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	zhangsan := &person{name: "张三", age: 20}
	//(*zhangsan).name = "张三2"
	zhangsan.name = "张三2"
	fmt.Println(zhangsan) //输出：&{张三2 20}

	wy := &[3]int{1, 2, 3}
	fmt.Println(wy[0])   //输出：1
	fmt.Println(wy[1:2]) //输出：[2]
}
