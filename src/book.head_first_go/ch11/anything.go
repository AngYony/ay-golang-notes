package main

import "fmt"

type Anything interface {
}

//定义一个类型
type SuperWhistle string

//为该类型定义方法
func (w SuperWhistle) Makesound() {
	fmt.Println("Makesound:", w)
}

//将空接口值作为参数
func Acceptanything(thing interface{}) {
	fmt.Println(thing)
	//使用类型断言来获得SuperWhistle
	whistle, ok := thing.(SuperWhistle)
	if ok {
		whistle.Makesound()
	}
}

func wyTest(wy Anything) {
	fmt.Println(wy)
}

func wyTest2(wy interface{}) {
	fmt.Println(wy)
}

func main() {
	Acceptanything(3.1415)
	Acceptanything(SuperWhistle("wwwwww"))

	wyTest2("wwweeeeeeeeeeeee")
	wyTest2(1111)
	wyTest2(false)
}
