package main

import "fmt"

type Toggleable interface {
	toggle()
}

type Switch string

//需要使用指针类型的接收器
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

func main() {

	s := Switch("off")
	s.toggle()

	var t Toggleable
	t = &s
	t.toggle()

}
