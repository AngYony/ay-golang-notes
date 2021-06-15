package main

import "fmt"

func say(worlds ...string) {
	//输出实参的类型
	fmt.Printf("%T", worlds) //输出：[]string
	fmt.Println(worlds)
}
func main() {
	say("aa", "bb", "cc")

	wy := []string{"AA", "BB"}
	fmt.Printf("%T", wy)

	//直接传入将会报编译错误
	//say(wy)
	say(wy...)

	var a = [...]int{1, 2}
	fmt.Println(a)
}
