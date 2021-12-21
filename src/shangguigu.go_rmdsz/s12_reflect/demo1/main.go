package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  uint8
}

// 基本数据类型
func reflectTest01(b interface{}) {
	// 通过反射获取变量的type,kind,value

	rTye := reflect.TypeOf(b)
	fmt.Printf("%T %v \n", rTye, rTye) // 输出：*reflect.rtype int

	rVal := reflect.ValueOf(b)         // 返回的类型是reflect.Value，不能直接参与变量的运算
	fmt.Printf("%T %v \n", rVal, rVal) // 输出：reflect.Value 100

	// 获取变量原来的值，并参与计算
	n := 10 + rVal.Int()
	fmt.Println("n=", n)

	// 将reflect.Value转回空接口类型
	iV := rVal.Interface()

	// 将空接口断言为具体类型，就可以参与运算
	num2, ok := iV.(int)
	if ok {

		fmt.Println("num2=", num2)
	}

}

// 反射结构体
func reflectStruct(v interface{}) {
	rTyp := reflect.TypeOf(v)
	fmt.Printf("%T %v \n", rTyp, rTyp) // 输出：*reflect.rtype main.Student

	rVal := reflect.ValueOf(v)
	fmt.Printf("%T %v \n", rVal, rVal) // 输出：reflect.Value {张三 11}

	fmt.Println("Kind:", rVal.Kind()) // 输出：struct

	iV := rVal.Interface()
	fmt.Printf("类型：%T ，值：%v \n", iV, iV) // 输出：类型：main.Student ，值：{张三 11}

	switch tv := iV.(type) {
	case Student:
		fmt.Println(tv.Name)
	}
}

func reflectUpdate(v interface{}) {
	rVal := reflect.ValueOf(v)
	// 如果实参是指针类型，那么此处的Kind将会输出 ptr
	fmt.Println("Kind：", rVal.Kind())
	rVal.Elem().SetInt(22)
	fmt.Println("通过Elem获取修改之后的值：", rVal.Elem().Int())
	// fmt.Println(rVal.Int())
}

func main() {
	var num int = 100
	reflectUpdate(&num)
	fmt.Println("修改之后的值：", num)

	// var num int = 100
	// reflectTest01(num)
	//
	// var stu = Student{
	// 	Name: "张三",
	// 	Age:  11,
	// }
	// reflectStruct(stu)
	// var a int = 1024
	// // 获取变量a的反射值对象(a的地址)
	// valueOfA := reflect.ValueOf(&a)
	// // 取出a地址的元素(a的值)
	// valueOfA = valueOfA.Elem()
	// // 修改a的值为1
	// valueOfA.SetInt(1)
	// // 打印a的值
	// fmt.Println(valueOfA.Int())
	//
	// var num int = 100
	// reflectTest01(num)

}
