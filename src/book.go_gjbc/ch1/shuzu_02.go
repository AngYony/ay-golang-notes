package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	// var b = &a //b是指向数组的指针
	// //通过数组指针访问数组元素的方式和通过数组类似
	// fmt.Println(b[0], b[1])
	// //通过for range来迭代数组指针指向的数组元素
	// for i, v := range b {
	// 	fmt.Println(i, v)
	// }

	for i := range a {
		fmt.Println(i, a[i])
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	var t [5][0]int
	for range t {
		fmt.Println("hello")
	}

	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		//struct{}部分是类型，{}表示对应的结构体值
		c2 <- struct{}{}
	}()
	<-c2

	fmt.Printf("%T", a)

}
