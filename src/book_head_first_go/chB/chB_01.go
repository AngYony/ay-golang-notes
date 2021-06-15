package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	asciiString := "abcde"
	utf8String := "中国人"
	//获取字节长度
	fmt.Println(len(asciiString))
	fmt.Println(len(utf8String))
	// 获取字符长度
	fmt.Println(utf8.RuneCountInString(asciiString))
	fmt.Println(utf8.RuneCountInString(utf8String))

	// asciiString := "abcde"
	// utf8String := "中国人"
	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	asciiPartial := asciiRunes[3:]
	utf8Partial := utf8Runes[2:]

	fmt.Println(string(asciiPartial)) //输出：de
	fmt.Println(string(utf8Partial))  //输出：人

}
