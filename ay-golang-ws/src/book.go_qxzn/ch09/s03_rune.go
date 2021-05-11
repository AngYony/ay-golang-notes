package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	zg := "中国"
	fmt.Println("字节长度：", len(zg))                    //输出：6
	fmt.Println("字符长度：", utf8.RuneCountInString(zg)) //输出：2

	c, size := utf8.DecodeRuneInString(zg)
	fmt.Printf("第一个字符：%c，其字节数为：%v \n", c, size) //输出：中和3

	for i, c := range zg {
		fmt.Printf("%v %c\n", i, c)
	}
}
