package main

import "fmt"

//定义一个新的类型
type MyType string

//函数被定义在MyType上，m表示一个接收器
func (m MyType) sayHi(wy string) {
	fmt.Println(m)
	fmt.Println(wy)
}

func main() {
	//创建一个MyType类型的值
	value := MyType("Hello")
	fmt.Println(value) //输出：Hello
	value.sayHi("good")

	// what := MyType("what")
	// fmt.Println(what)
	// what.sayHi("are")

}
