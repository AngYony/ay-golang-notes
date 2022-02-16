package main

import "fmt"

func main() {
	var name string = "s"

	a := []int{1, 2, 3}
	b := a
	c := a[1:3]
	d := a[1]
	a[1] = 4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
