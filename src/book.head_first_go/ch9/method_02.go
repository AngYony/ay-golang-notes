package main

import "fmt"

type Number int

//值类型传递
func (n Number) Double() {
	n *= 2
}

//指针类型传递值
func (n *Number) PointerDouble() {
	*n *= 2
}

func main() {
	mynum := Number(4)
	pointer := &mynum     //指针
	mynum.PointerDouble() //值类型mynum自动转换为指针，调用指针类型方法
	fmt.Println(mynum)    //输出：8，同时指向该变量的指针对应的值都为8

	pointer.Double() //指针自动转换为值类型
	pointer.PointerDouble()
	fmt.Println(*pointer) //输出16

	a := Number(4).PointerDouble()
}
