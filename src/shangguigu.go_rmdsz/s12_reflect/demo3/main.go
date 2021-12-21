package main

import (
	"fmt"
	"reflect"
)

type user struct {
	UserId string
	Name   string
}

func main() {
	var (
		model *user // user的指针
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)                               // 获取类型 *user
	fmt.Println("reflect.TypeOf.Kind =", st.Kind().String()) // ptr

	st = st.Elem()                                           // st指向的类型
	fmt.Println("reflect.TypeOf.Elem =", st.Kind().String()) // struct

	// New返回一个Value类型值，该值持有一个指向类型为Type的新申请的零值的指针
	elem = reflect.New(st)                                              // 返回一个指向User的指针
	fmt.Println("reflect.New.Kind =", elem.Kind().String())             // ptr
	fmt.Println("reflect.New.Elem.Kind =", elem.Elem().Kind().String()) // struct

	// model就是创建的user结构体变量，先转换为空接口，再断言一下
	model = elem.Interface().(*user) // model是*user，它的指向和elem是一样的
	elem = elem.Elem()               // 取得elem指向的值
	elem.FieldByName("UserId").SetString("1111111")
	elem.FieldByName("Name").SetString("展示")
	fmt.Println(model)
}
