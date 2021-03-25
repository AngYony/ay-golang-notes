package main

import "fmt"

//定义一个名为part的类型
type part struct {
	desc  string
	count int
}

//定义一个名为car的类型
type car struct {
	name     string
	topSpeed float64
}

func run(c *car) {
	c.name = "wt"
	c.topSpeed = 11.1
	(*c).name = "abc"
}
func main() {
	var myCar car //定义一个car类型的变量
	myCar.name = "hi"
	myCar.topSpeed = 30.0
	fmt.Println(myCar) //输出：{hi 30}
	run(&myCar)
	fmt.Println(myCar) //输出：{wt 11.1}
}
