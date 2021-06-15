package main

import "fmt"

func main() {
	type wy struct {
		h int
		w int
	}

	a := wy{1, 2}
	b := a
	b.h += 5
	fmt.Println(b) //输出：{6 2}
	fmt.Println(a) //输出：{1 2}

	c := []wy{
		{h: 3, w: 1},
		{h: 2, w: 3},
		{h: 4, w: 5}, //末尾必须有逗号
	}
	fmt.Println(c)

}
