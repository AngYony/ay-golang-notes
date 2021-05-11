package main

import (
	"fmt"
	"strconv"
)

func main() {

	var pi rune = 'π'
	//不进行转换，将会按照int32类型输出值
	fmt.Println(pi) //输出：960
	//转换为字符串
	fmt.Println(string(pi)) //输出：π

	//此处必须是双引号，不能是单引号
	var pi2 string = "π"
	fmt.Println([]rune(pi2))

	countdown := 65
	//输出该代码点表示的字符
	fmt.Printf("%c \n", countdown) //输出：A
	//将整数转换为ASCII字符
	str := strconv.Itoa(countdown) //输出：65
	fmt.Println(str)
	//将数值转换为字符串
	str = fmt.Sprintf("%v", countdown) //输出：65
	fmt.Println(str)

	c, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("转换错误")
	}
	fmt.Println(c) //输出：10
}
