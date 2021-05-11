package main

import "fmt"

func main() {

	// year := 2021
	// fmt.Printf("%T\n", year)
	// fmt.Printf("%v\n", year)

	// a := []int{1, 2, 3}
	// fmt.Printf("%T", a)

	// fmt.Printf("%v", a)

	// year = 12
	// fmt.Printf("%v 的类型为：%[1]T  \n", year)

	// var flag = true
	// fmt.Printf("%T", flag)

	//定义十进制的数字
	var red, green, blue uint8 = 0, 141, 213
	//转换为16进制显示
	fmt.Printf("%x %x %x", red, green, blue) //输出：0 8d d5
	//定义十六进制的数字
	var red2, green2, blue2 uint8 = 0x0, 0x8d, 0xd5
	//转换为10进制数
	fmt.Printf("%d %d %d", red2, green2, blue2) //输出：0 141 213
	//笔记不同进制的值
	fmt.Println(red == red2) //输出：true

	fmt.Printf("%02x %02x %02x;", red, green, blue)

	// fmt.Printf("color:#%02x%02x%02x;", red, green, blue)

	fmt.Printf("%b", green)
}
