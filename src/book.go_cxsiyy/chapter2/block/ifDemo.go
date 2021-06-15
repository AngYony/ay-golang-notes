package main

import (
	"fmt"
	"os"
)

func f() int {
	return 1
}
func g(n int) int {
	return n + 1
}
func main() {
	if x := f(); x == 0 { //定义变量x，值为1
		fmt.Println(x)
	} else if y := g(x); x == y { //定义变量y，值为x+1=2
		fmt.Println(x, y) //如果满足条件时：输出1,1
	} else {
		fmt.Println(x, y) //输出：1，2
	}
	//fmt.Println(x, y) //编译错误
}

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		print(err)
	}
}
