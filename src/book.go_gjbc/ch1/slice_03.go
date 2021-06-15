package main

import "fmt"

func main() {
	var (
		a []int               // nil切片，和nil相等，一般用来表示一个不存在的切片
		b = []int{}           // 空切片，和nil不相等，一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片，len和cap都为3
		d = c[:2]             // 有2个元素的切片，len为2，cap为3
		e = c[0:2:cap(c)]     // 有2个元素切片，len为2，cap为3
		f = c[:0]             // 有0个元素的切片，len为0，cap为3
		g = make([]int, 3)    // 有3个元素的切片，len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片，len为2，cap为3
		i = make([]int, 0, 3) // 有0个元素的切片，len为0，cap为3

	)
	fmt.Printf("a：%#v\n", a)
	fmt.Printf("b：%#v\n", b)
	fmt.Printf("c：%#v\n", c)
	fmt.Printf("d：%#v\n", d)
	fmt.Printf("e：%#v\n", e)
	fmt.Printf("f：%#v\n", f)
	fmt.Printf("g：%#v\n", g)
	fmt.Printf("h：%#v\n", h, cap(h))
	fmt.Printf("i：%#v\n", i)

}
