package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

//形参是接口变量
func shout(t talker) {
	fmt.Println(t.talk())
}

type martian struct{}

//实现talker接口，注意接收者是类型本身
func (m martian) talk() string {
	return "kao"
}

type laser int

//实现talker接口方法
func (l *laser) talk() string {
	return strings.Repeat("A", int(*l))
}

func main() {
	shout(martian{})
	shout(&martian{})
	pew := laser(2)
	shout(&pew)
	//运行将会报错
	shout(pew)
}
