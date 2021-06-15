package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "大"
}

type laser int

func (l laser) talk() string {
	//重复输出指定个数的字符
	return strings.Repeat("小", int(l))
}

func shout(t talker) {
	fmt.Println(t.talk())
}

func main() {

	var t1 talker = martian{}
	fmt.Println(t1.talk())

	var t2 talker = laser(3)
	fmt.Println(t2.talk())
}
