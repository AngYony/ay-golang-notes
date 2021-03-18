package main

import (
	"fmt"
)

func main() {
	amount := 6
	double(&amount)   //传递一个指针
	fmt.Print(amount) //输出12

	truth := true
	negate(&truth)
	fmt.Println(truth)

}

//函数的参数是指针类型
func double(number *int) {
	*number *= 2 //设置指针number处的值乘以2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
