package main

import (
	"fmt"
)

func main() {
	a := "A"
	wy := &a         //获取指向a的指针wy
	fmt.Println(*wy) //获取指针指向的变量的值，输出：A

	b := *wy         //将指针指向的变量的值赋值给另一个变量
	*wy = "B"        //修改原来的指针指向的变量的值
	fmt.Println(b)   //输出：A
	fmt.Println(*wy) //输出：B

	c := 1
	d := c
	c = 3
	fmt.Println(d, c) //输出CN

}
