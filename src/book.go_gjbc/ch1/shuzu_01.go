package main

import "fmt"

func main() {
	var a [3]int

	var b = [...]int{1, 2, 3}
	var b2 = [3]int{4, 5, 6}
	var c = [...]int{2: 3, 1: 2}
	var d = [...]int{1, 2, 4: 5, 6}

	fmt.Println(a) //输出：[0 0 0]
	fmt.Println(b) //输出：[1 2 3]
	fmt.Println(b2)
	fmt.Println(c) //输出：[0 2 3]
	fmt.Println(d) //输出：[1 2 0 0 5 6]
}
