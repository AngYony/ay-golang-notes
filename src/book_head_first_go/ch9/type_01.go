package main

import "fmt"

//基于float64定义一个新类型
type Gallons float64
type Liters float64

func main() {
	var carFuel Gallons     //定义一个Gallons类型的变量
	carFuel = Gallons(10.0) //把float64转换为Gallons
	myLiter := Liters(11.1)
	fmt.Println(carFuel, myLiter)

}
