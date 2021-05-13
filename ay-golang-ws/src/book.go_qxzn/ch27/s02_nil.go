package main

import "fmt"

type number struct {
	value int
	valid bool //零值为false
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

//实现Stringer接口
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)

}

func main() {
	n := newNumber(33)
	fmt.Println(n) //执行String()方法，输出value的值：33
	e := number{}
	fmt.Println(e) //输出：not set

}
