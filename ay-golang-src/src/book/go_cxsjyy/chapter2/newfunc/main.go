package main

import "fmt"

func main() {
	p := new(int)   //定义类型是整型指针（*int）的p，指针p指向未命名的int变量
	fmt.Println(p)  //输出：0xc00005c080
	fmt.Println(*p) //输出：0
	*p = 2          //把未命名的int设置为2
	fmt.Println(*p) //输出2

	a := new(int)
	b := new(int)
	fmt.Println(a) //0xc00005c0b0
	fmt.Println(b) //0xc00005c0b8

}

func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dumy int
	return &dumy
}

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
