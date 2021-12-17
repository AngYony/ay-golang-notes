# Go - reflect

反射可以在运行时动态获取变量的各种信息，比如变量的类型（type），类别（kind）。

如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）。

==通过反射，可以修改变量的值，可以调用关联的方法。==



## reflect.Type

reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型。

通过reflect.Type.Kind也可以获取Kind值，Kind代表Type类型值表示的具体分类，返回的是一个常量。

Type和Kind:

- Type是变量的类型，Kind是变量的类别，Type 和 Kind 可能是相同的，也可能是不同的。
- `var stu Student`中，stu的type是包名.Student，而Kind是struct；`var num int=10`中，num的Type和Kind都是int。





## reflect.Value

reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型。

reflect.Value是一个结构体类型，可以获取到关于该变量的很多信息。

在反射中，变量、interface{} 和 reflect.Value 是可以相互转换的，会高频使用到。

reflect.Value.Kind，获取变量的类别，Kind代表Type类型值表示的具体分类，返回的是一个常量。注意：它不是变量的类型，返回的更像是定义变量类型时使用到的关键字。例如，数组类型的变量的Kind值是Array，切片类型的变量的Kind值是Map。



### reflect.Value.Elem

Elem返回reflect.Value持有的接口保管的值的封装，或者指针指向的值的封装。如果变量的Kind不是interface或Ptr会panic，如果变量的值为nil，会返回变量零值。

Elem一般用于通过反射来修改变量值的操作。

```go
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
}
```

rVal.Elem().SetInt(22) 等价于：

```go
var num = 10
var b *int = &num
*b = 22  // 为指针指向的变量赋值
```

rVal.Elem()等价于`*b`。



## 通过反射修改值

通过反射来修改变量时，需要注意：当使用SetXxx方法来设置值时，需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用到reflect.Value.Elem()方法。



## 通过反射操作结构体

使用反射来遍历结构体的字段，并调用结构体的方法，和获取结构体标签的值。

### 遍历结构体字段

获取字段的个数：

```go
func (v Value) NumField() int 
```

获取字段的值：

```go
func (v Value) Field(i int) Value 
```

该方法返回的是reflect.Value，只能显示值，而不能直接参与运算。

获取字段本身：

```go
Field(i int) StructField
```

获取字段的标签：

```go
func (tag StructTag) Get(key string) string
```

必须通过reflect.Type.Field(i)来获取Tag。

```go
tagVal := typ.Field(i).Tag.Get("json")
```



具体使用，见下文的示例。



### 调用结构体方法

获取方法条目数：

```
NumMethod() int
```

注意：

- 当反射的变量是结构体指针时，无论结构体中的方法在定义的时候，接收器指定的是否是struct指针，都可以获取到所有方法的数量。
- 当反射的变量是结构体（非指针时）时，只能获取结构体中接收器是struct的方法数量，接收器为struct指针的方法不能获取到。

具体见综合示例中的该方法注释部分。

获取方法：

```go
func (v Value) Method(i int) Value
```

默认按方法名排序对应i值，i从0开始。

调用方法：

```go
func (v Value) Call(in []Value) []Value
```

传入参数和返回参数都是eflect.Value类型的数组。

 综合示例一，获取结构体的字段、标签，并修改字段的值和调用结构体的方法。

```go
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

```

输出：

```
*main.Student &{张三 20} ptr
该结构体一共有4个方法           
该结构体一共有2个字段           
字段Name的值为：张三            
字段0的标签为：my_name          
字段Age的值为：20               
字段1的标签为：my_age           
我的名字是： 张三               
调用第1个方法，执行的结果为： 30
修改字段后的结果： {李四 20}  
```



示例二，使用反射实例化结构体。

```go
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
```

输出结果：

```
reflect.TypeOf.Kind = ptr
reflect.TypeOf.Elem = struct  
reflect.New.Kind = ptr        
reflect.New.Elem.Kind = struct
&{1111111 展示}  
```





## 通过反射获取值

```go
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
	n := 10 + rVal.Int() // 重点
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

func main() {
	
	var num int=100
	reflectTest01(num)
	
	var stu = Student{
		Name: "张三",
		Age:  11,
	}
	reflectStruct(stu)

}

```



