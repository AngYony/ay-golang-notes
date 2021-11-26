package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 3.0
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer) //输出指针处的值
}

func main() {

	var myBool bool = true
	printPointer(&myBool)

	//var myFloatPointer2 *float64 = createPointer()
	//fmt.Println(*myFloatPointer2)

	//var myFloat float64
	//fmt.Println(&myFloat) //获取float64类型的变量的地址

	//var myBool bool
	//fmt.Println(&myBool) //获取bool类型的变量的地址

	//fmt.Println(reflect.TypeOf(&myFloat))
	//fmt.Println(reflect.TypeOf(&myBool))

	//var myInt int = 4
	//fmt.Println(&myInt)                 //获取int类型的变量的地址（获取指针）
	//fmt.Println(reflect.TypeOf(&myInt)) //获取指向myInt的指针的类型（获取指针的类型）
	//
	//var myIntPointer *int      //声明一个指向int的指针变量（声明*int类型的指针变量）
	//myIntPointer = &myInt      //为指针变量分配一个指针（为指针变量赋值，值必须是一个指向相同类型的指针）
	//fmt.Println(myIntPointer)  //打印指针本身
	//fmt.Println(*myIntPointer) //打印指针处的值
	//
	//*myIntPointer = 8
	//fmt.Println(myInt)
	//
	//myFloat := 1.1
	//fmt.Println(myFloat)
	//myFloatPointer := &myFloat   //定义一个指针（值为变量myFloat的地址）
	//*myFloatPointer = 2.2        //给指针处的变量赋一个值（更新地址上的值)
	//fmt.Println(*myFloatPointer) //打印指针对应的地址上的值
	//fmt.Println(myFloat)         //指向该地址的变量的值也一并发生改变

}
