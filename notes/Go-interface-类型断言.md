# Go - 接口与类型断言

Go中接口的用途和C#中的接口类似。

一个接口是特定值预期具有的一组方法。一个接口类型是一个抽象类型。接口并不会描述值是什么：它们不会说基础类型是什么或者数据是如何保存的。它们仅仅描述值能做什么：它具有哪些方法。



## 定义接口

使用interface关键字定义一个接口类型，后面跟着一个花括号，内部含有一组方法，以及方法期望的参数和返回值。

为了便于复用，通常会把接口声明为类型并为其命名。按照惯例，接口类型的名称常常会以-er作为后缀。

```go
type Player interface {
	Play(string)
	Stop()
}
```

接口类型的规则和其他类型相同。如果名称以小写字母开头，接口类型不会被导出并且不会在当前包之外被访问。



## 实现接口

在C#中，实现一个接口需要显式指定，而在Go语言中，这是自动发生的。（接口满足是自动的。不需要显式声明具体类型满足Go中的接口。）

实现接口的类型，需要满足的条件：该类型必须实现接口中定义的所有方法，并且这些方法名称、参数类型（可能没有）和返回值（可能没有）都必须和接口中定义的一致。

任何拥有接口定义的所有方法的类型被称作满足那个接口。

除了接口中列出的方法之外，类型还可以有更多的方法，但是它不能缺少接口中的任何方法，否则就不满足那个接口。

一个类型可以满足多个接口，一个接口（通常应该）可以有多个类型满足它。

```go
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
```

如果一个类型包含接口中声明的所有方法，那么它可以在任何需要接口的地方使用，而不需要更多的声明。



## 接口类型的变量

一个接口类型的变量能够保存任何满足接口的类型的值。

```go
import (
	"book_head_first_go/ch11/mypkg"
	"fmt"
)

func main() {
	//声明一个接口类型的变量
	var value mypkg.MyInterface
	//将MyType类型的值赋值给该接口变量（MyType满足MyInterface接口）
	value = mypkg.MyType(5)
	//调用满足该接口的对应的方法
	value.MethodWithoutParameters()
	value.MethodWithParameter(11.1)
	fmt.Println(value.MethodWithReturnValue())
}
```

也可以将函数的参数定义为接口类型。

```go
type NoiseMaker interface {
	MakeSound()
}
func play(n NoiseMaker) {
	n.MakeSound()
}
```

一旦你给一个接口类型的变量（或方法的参数）赋值，你就只能调用接口定义的方法。(和C#类似)

只要是满足接口的类型，即使该类型具有其他方法，都可以将该类型赋值给该接口变量。

### 将指针传递给接口变量

如果一个类型声明了指针接收器方法，就只能将那个类型的指针传递给接口变量。

```go
type Toggleable interface {
	toggle()
}

type Switch string

//需要使用指针类型的接收器
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}
```

调用代码：

```go
var t Toggleable	//定义接口变量
t = &s				//将指针赋值给接口变量
t.toggle()
```

【也可参考指针笔记中的“接口与指针”部分】



## 类型断言

当有一个接口类型的变量时，你能在它之上调用的方法只能是接口定义的。如果你将一个具体类型赋值给接口类型的变量，你能使用类型断言来获得具体类型的值。只有这样你才能调用具体类型定义的（但没有定义在接口中的）方法。

当你将一个具体类型的值赋给一个接口类型的变量时，类型断言让你能取回具体类型。有点类似C#中的类型转换和is关键字。

假如存在如下接口：

```go
type NoiseMaker interface {
	MakeSound()
}
```

Robat类型实现了该接口：

```go
type Robat string

func (r Robat) MakeSound() {
	fmt.Println("Robat_MakeSound:", r)
}

//定义一个额外的方法
func (r Robat) Walk() {
	fmt.Println("Robat_Walk:", r)
}
```

将具体值赋值给接口变量：

```go
var noiseMaker NoiseMaker = Robat("FFF")
noiseMaker.MakeSound()
```

**使用类型断言，获取具体类型**，并调用接口中没有的方法：

```go
var robot Robat = noiseMaker.(Robat)
robot.Walk()
```

简单来说，类型断言就像说某物像“我知道这个变量使用接口类型NoiseMaker，但是我很确信这个NoiseMaker实际上是Robot。”

一旦你使用类型断言来取回具体类型的值，你可以调用那个类型上的方法，包括接口中未声明的方法。

类型断言返回第二个bool值来表明断言是否成功。

类型断言失败时，为了避免异常，可以使用断言返回的第二个可选值：

```go
var wy, ok = noiseMaker.(Robat)
if ok {
	wy.Walk()
} else {
	fmt.Println("失败")
}
```

> 如果类型断言被用于期待多个返回值的情况，它能有第二个可选的返回值来表明断言是否成功。（并且断言并不会在不成功时出现异常。）第二个值是一个bool，并且当原类型和断言类型相同时，返回true，否则返回false。你可以对于第二个返回值做任何操作，但是按照惯例，它通常被赋给一个名为ok的变量。
>
> ```go
> recorder, ok := player.(gadget.TapeRecorder)
> if ok {
> 	recorder.Record()
> } else {
> 	fmt.Println("断言失败")
> }
> ```



## error 接口

error类型本质上只是一个接口：

```go
type error interface{
	Error() string
}
```

因此，如果自定义的类型中，包含一个返回string的Error方法，那么它就满足error接口，并且它是error的值。

自定义error类型：

```go
//定义一个以string为基础类型的类型
type ComedyError string
//满足error接口
func (c ComedyError) Error() string {
	return string(c)
}

//定义一个基础类型为float64的类型
type OverheatError float64
//实现error接口
func (o OverheatError) Error() string {
	return fmt.Sprintf("值：%0.2f", o)
}
//指定函数返回原生error值
func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
        //var wyerr error
        //wyerr = OverheatError(excess)
        //return wyerr
		return OverheatError(excess) //返回的是error，实现的具体类型
	}
	return nil
}

func main() {
	var err error //声明一个error类型的变量
	err = ComedyError("这是一个错误信息")	//ComedyError满足error接口，所以可以赋值给接口变量
	fmt.Println(err)

	err = checkTemperature(121.322, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
```

> error类型像int或者string一样是一个“预定义标识符”，它不属于任何包。它是“全局块”的一部分，这意味着它在任何地方可用，不用考虑当前包信息。



## Stringer 接口

fmt包中定义了fmt.Stringer接口：

```go
type Stringer interface{
	String() string
}
```

因此，任何具有返回string的String()方法的类型都是一个fmt.Stringer。

```go
type CoffeePot string

//满足Stringer接口
func (c CoffeePot) String() string {
	return string(c) + "浪浪浪"
}
```

许多在fmt包中的函数都会判断传入的参数是否满足stringer接口，如果满足就调用String方法。这些函数包括Print、Println和Printf等。

```go
func main() {
	coffeePot := CoffeePot("wogan")
	fmt.Println(coffeePot.String())
	fmt.Println(coffeePot)
	fmt.Printf("%s", coffeePot)
}
```

该功能一般应用在格式输出的场景，或自定义格式字符串。



## 空接口

如果接口声明定义了方法，类型必须实现这个方法才能满足接口。如果定义了一个不需要任何方法的接口，它会被任何类型满足！它被所有的类型满足！

```go
//定义一个空接口
type Anything interface {
}
```

空接口，可以用来接收任何类型的值。

如果你定义一个接收一个空接口作为参数的函数，你可以传入任何类型的值作为参数。

```go
//定义一个参数为空接口的函数
func wyTest(wy Anything) {
	fmt.Println(wy)
}

func main() {
    //可以为该函数指定任意类型的实参
	wyTest("wwweeeeeeeeeeeee")
	wyTest(1111)
	wyTest(false)
}
```

除了上述定义空接口的方式外，常用的是直接使用下述形式：

```go
func wyTest2(wy interface{}) {
	fmt.Println(wy)
}
```

不需要实现任何方法来满足空接口，所以所有的类型都满足它。

### 使用断言在空接口值上调用函数

如果你有一个接口类型的值，你只能调用接口上的方法。空接口没有任何方法。那意味着你没法调用空接口类型值的任何方法。

为了在空接口类型的值上调用方法，你需要使用类型断言来获得具体类型的值。

```go
//定义一个类型
type SuperWhistle string

//为该类型定义方法
func (w SuperWhistle) Makesound() {
	fmt.Println("Makesound:", w)
}

//将空接口值作为参数
func Acceptanything(thing interface{}) {
	fmt.Println(thing)
	//使用类型断言来获得SuperWhistle
	whistle, ok := thing.(SuperWhistle)
	if ok {
		whistle.Makesound()
	}
}
```

调用方式：

```go
func main() {
	Acceptanything(3.1415)
    //传入SuperWhistle值
	Acceptanything(SuperWhistle("wwwwww"))
}
```

