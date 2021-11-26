package main

import "fmt"

type CoffeePot string

//满足Stringer接口
func (c CoffeePot) String() string {
	return string(c) + "浪浪浪"
}

func main() {
	coffeePot := CoffeePot("wogan")
	fmt.Println(coffeePot.String())
	fmt.Println(coffeePot)
	fmt.Printf("%s", coffeePot)
}
