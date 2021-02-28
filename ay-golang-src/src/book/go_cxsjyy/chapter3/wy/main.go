package main

import (
	"fmt"
	"math"
)

func main() {
	var i int8 = 127
	fmt.Println(i + 1) //结果为：-128
	fmt.Println(i * i) //结果为：1

	var u uint8 = 255
	fmt.Println(u + 1) //结果为：0
	fmt.Println(u * u) //结果为：1

	var u2 uint8 = 0
	fmt.Println(u2 - 1) //结果为：255

	o := 0666                           //定义一个八进制数
	fmt.Printf("%d \n", o)              //输出：438
	fmt.Printf("%o \n", o)              //输出：666
	fmt.Printf("%#o \n", o)             //按照标准前缀，输出：0666
	fmt.Printf("%d %[1]o %#[1]o \n", o) //输出：438 666 0666

	fmt.Printf("%x \n", o)  //输出：1b6
	fmt.Printf("%X \n", o)  //输出：1B6
	fmt.Printf("%#x \n", o) //输出：0x1b6

	x := int64(0xdeadbeef)                   //定义十六进制数
	fmt.Printf("%d %[1]x %#[1]x %#[1]X ", x) // 输出：3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	fmt.Println(math.MaxFloat64)

	var f1 float32 = 999999
	fmt.Println(f1) //输出：999999

	var f2 float32 = 999999 + 1
	fmt.Println(f2) //输出：1e+06

	var f3 float32 = 16777215 //临界值
	fmt.Println(f3+1)   //输出：1.6777216e+07
	fmt.Println(f3 == f3+1)   //输出：true

	for x := 0; x < 8; x++ {
		//fmt.Println(math.Exp(float64(x))) //输出自然对数e的各个幂方值
		fmt.Printf("x=%d e^x=%8.3f \n", x, math.Exp(float64(x)))
	}

	fmt.Print(1<<4)






}
