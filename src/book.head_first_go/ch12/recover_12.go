package main

import "fmt"

func calmDown2() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown2()
	err := fmt.Errorf("错误消息")
	panic(err)
}
