package main

import "fmt"

type person struct {
	name string
	age  int
}

//使用指针作为接收者
func (p *person) birthday2() {
	p.age++
}

func birthday(p *person) {
	p.age++
}

func main() {

	zhangsan := person{name: "张三", age: 22}
	birthday(&zhangsan)
	fmt.Println(zhangsan)
	zhangsan.birthday2()

	//声明一个指针类型变量
	terry := &person{name: "关羽", age: 33}
	terry.birthday2()
	fmt.Println(*terry)

}
