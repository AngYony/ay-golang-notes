package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[i]=%c\n", x[i])
		x := x[i] //重新什么一个局部变量x
		if x != '!' {
			x := x + 'A' - 'a' //再次声明一个局部变量x
			fmt.Printf("%c\n", x)
		}
	}
}
