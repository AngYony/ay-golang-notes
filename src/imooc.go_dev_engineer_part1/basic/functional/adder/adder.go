package main

import "fmt"

func adder() func(int) int {
	sum := 0
	//不仅返回函数，还返回对变量sum的引用
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder() //a的函数体中存有sum

	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}
}
