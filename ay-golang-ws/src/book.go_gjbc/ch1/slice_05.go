package main

import "fmt"

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func wyFun(x byte) bool {
	return x == '-'
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:0]
	fmt.Println(len(b), cap(b)) //输出：0和5

	wy := Filter([]byte("a-b-c"),
		func(x byte) bool {
			return x == '-'
		})
	fmt.Println(string(wy)) //输出：abc
}
