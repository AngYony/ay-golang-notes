package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

type hello struct {
	car //匿名字段，其类型为car本身
	str string
}

func main() {
	//创建struct类型car的变量myCar，并同时为字段赋值
	myCar := car{name: "hi", topSpeed: 20}
	//为struct类型字段赋值
	hi := hello{car: myCar, str: "woqu"}
	fmt.Println(hi) //输出：{{hi 20} woqu}

	//获取嵌入外部struct的字段名称
	fmt.Println(hi.name)
	fmt.Println(hi.topSpeed)

}
