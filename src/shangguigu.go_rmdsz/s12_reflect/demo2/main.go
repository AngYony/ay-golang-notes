package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"my_name"`
	Age  uint8  `json:"my_age"`
}

func (s *Student) SayName() {
	fmt.Println("我的名字是：", s.Name)
}

func (s *Student) AddAge(age uint8) {
	s.Age += age
}
func (s *Student) GetSum(n1, n2 int) int {
	return n2 + n1
}

func (s Student) SayNo() {
	fmt.Println("NONONO")
}

func reflectStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Println(typ, val, kd)

	// 判断是否是结构体，引入目标常量来判断，只有是struct或者struct的指针可以解析
	if kd != reflect.Struct && kd != reflect.Ptr {
		fmt.Println("不能解析结构体")
		return
	}

	// 获取结构体的方法个数
	/*
		1.当反射的变量是结构体指针时，无论结构体中的方法在定义的时候，接收器指定的是否是struct指针，都可以获取到
		2.当反射的变量是结构体（非指针时）时，只能获取结构体中接收器是struct的方法数量，接收器为struct指针的方法不能获取到
	*/
	numOfMethod := typ.NumMethod()
	fmt.Printf("该结构体一共有%d个方法\n", numOfMethod)

	// 获取结构体字段数目
	var num int
	if kd == reflect.Struct {
		num = typ.NumField()

	} else if kd == reflect.Ptr {
		num = typ.Elem().NumField()
	}
	fmt.Printf("该结构体一共有%d个字段\n", num)

	if kd == reflect.Struct {
		// 遍历结构体所有字段
		for i := 0; i < num; i++ {
			// 获取字段的值，val.Field(i)不能直接进行运算
			fmt.Printf("字段%s的值为：%v\n", typ.Field(i).Name, val.Field(i))
			// 获取字段的标签
			tagVal := typ.Field(i).Tag.Get("json")

			if tagVal != "" {
				fmt.Printf("字段%d的标签为：%v\n", i, tagVal)
			}
		}

	} else if kd == reflect.Ptr {
		// 遍历结构体所有字段
		for i := 0; i < num; i++ {

			// 获取字段的值，val.Field(i)不能直接进行运算
			fmt.Printf("字段%s的值为：%v\n", typ.Elem().Field(i).Name, val.Elem().Field(i))
			// 获取字段的标签
			tagVal := typ.Elem().Field(i).Tag.Get("json")
			if tagVal != "" {
				fmt.Printf("字段%d的标签为：%v\n", i, tagVal)
			}
		}
	}

	if kd == reflect.Struct {
		// 由于结构体中接收器是struct的只有一个方法，因此当传入的变量是struct时，只能获取一个方法序号
		val.Method(0).Call(nil)
	} else if kd == reflect.Ptr {

		// 能够获取到所有方法

		// 调用没有参数的方法
		// 按照函数名称排序，从0开始，调用第3个方法
		val.Method(2).Call(nil) // 无参数函数调用

		// 调用有参数的方法
		// 什么一个存放参数的切片
		var params []reflect.Value
		params = append(params, reflect.ValueOf(10))
		params = append(params, reflect.ValueOf(20))
		// 调用 GetSum方法，传入两个参数
		res := val.Method(1).Call(params)
		fmt.Println("调用第1个方法，执行的结果为：", res[0].Int())

		// 修改字段的值
		val.Elem().FieldByName("Name").SetString("李四")

	}
}
func main() {
	stu := Student{
		Name: "张三",
		Age:  20,
	}

	reflectStruct(&stu)
	fmt.Println("修改字段后的结果：", stu)
}
