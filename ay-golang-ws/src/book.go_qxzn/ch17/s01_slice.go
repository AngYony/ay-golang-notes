package main

import "fmt"

func main() {

	wy := "abcde"
	abc := wy[:3]
	//为变量重新赋值并不会改变已切分得到的新字符串
	wy = "efgh"
	fmt.Println(abc) //仍然输出：abc

	zg := "中国"
	s := zg[:3]
	fmt.Println(s) //输出：中

}
