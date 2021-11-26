package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

//定义一个指定传入的函数类型
func doMath(mathFunction func(int, int) float64) {
	result := mathFunction(10, 2)
	fmt.Println(result)
}
func main() {
	var greeterFunction func()
	greeterFunction = sayHi
	greeterFunction()

	//定义函数类型变量
	var mathFunction func(int, int) float64
	mathFunction = divide //为变量分配函数
	//将函数类型变量作为参数传递给以函数作为形参的函数，类似于C#委托
	doMath(mathFunction)

}
