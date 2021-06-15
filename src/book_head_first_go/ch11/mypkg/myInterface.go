package mypkg

import "fmt"

//声明一个接口类型
type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

//定义MyType类型，实现MyInterface接口
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters被调用")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "MethodWithReturnValue被调用"
}

//一个类型即使有额外的不属于接口的方法，但它仍然可以满足接口
func (m MyType) MethodNotInInterface() {
	fmt.Println("该方法不在接口方法内")
}
