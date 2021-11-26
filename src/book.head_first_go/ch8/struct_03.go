package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

//方式一（不推荐）：直接返回struct类型，适用于小型struct
func createCar(name string, topSpeed float64) car {
	var oneCar car
	oneCar.name = name
	oneCar.topSpeed = topSpeed
	return oneCar
}

//方式二（推荐）：返回一个struct类型的指针
func createCar2(name string, topSpeed float64) *car {
	var oneCar car
	oneCar.name = name
	oneCar.topSpeed = topSpeed
	return &oneCar
}

//修改操作
func editCar(c *car, name string, topSpeed float64) {

	c.name = name
	c.topSpeed = topSpeed
}

//获取指针变量的值
func getCar(c *car) {
	fmt.Println("Name:", c.name)
	fmt.Println("TopSpeed:", (*c).topSpeed)
}

func main() {
	//方式一的调用
	myCar := createCar("wy", 11.1)
	fmt.Println(myCar) //输出：{wy 11.1}

	//方式二的调用
	myCar2 := createCar2("aa", 22.2)
	fmt.Println(myCar2) //输出：&{aa 22.2}

	//修改操作的调用
	editCar(myCar2, "bb", 33.3)
	fmt.Println(myCar2) //输出：&{bb 33.3}

	//获取指针变量的值
	getCar(&myCar)
	getCar(myCar2)
}
