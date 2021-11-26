package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//去掉空格
	fmt.Println(strings.TrimSpace(" wy ") + "ang")

	//将字符串转换为float64
	f, err := strconv.ParseFloat("12.34", 64)
	if err != nil {
		fmt.Println("有错误")
	}
	if f >= 12 {
		fmt.Println("大于12")
	}

	y := 39.995
	fmt.Println(strconv.FormatFloat(y+0.005, 'f', 2, 64))

}
