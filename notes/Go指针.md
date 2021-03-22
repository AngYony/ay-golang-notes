# Go指针

指针：表示变量地址的值称为指针，指针指向变量的位置（地址）。

> 指针的值是一个变量的地址（可以直接将指针理解为变量的地址），因此，一个指针指示了变量的值所保存的位置。不是所有的值都有地址，但是所有的变量都有。使用指针，可以在无须知道变量名字的情况下，间接读取或更新变量的值。



## 指针的值的表示方式：地址运算符（&变量名）

地址运算符：`&变量名`（可以使用一个`&`符号获取变量的地址）。

> 例如，声明了一个变量x：
>
> ```go
> var x int
> ```
>
> 获取变量x的地址，可以使用表达式：&x，&x表示的是获取一个指向整型变量x的指针。因此&x就代表着指针本身，它的值是一个变量的地址。在这里，它的类型是整型指针（*int）。

示例，获取任意类型变量的地址：

```go
func main() {

	var myInt int
	fmt.Println(&myInt) //获取int类型的变量的地址

	var myFloat float64
	fmt.Println(&myFloat) //获取float64类型的变量的地址

	var myBool bool
	fmt.Println(&myBool) //获取bool类型的变量的地址
}
```

运行后，输出内容如下：

```
API server listening at: 127.0.0.1:46048
0xc000014098
0xc0000140d0
0xc0000140d8
Process exiting with code: 0
```

总结：指针表示的是变量的地址，使用`&变量名`的形式来表示指针的值。



## 指针的类型的表示方式：*类型

指针的类型可以写为一个*符号，后面跟着指针指向的变量的类型。

例如，指向一个`int`变量的指针的类型将被写为`*int`，读作“==指向int的指针==“。

示例，使用reflect.TypeOf函数来显示之前程序中指针的类型：

```go
fmt.Println(reflect.TypeOf(&myInt)) //获取指向myInt的指针的类型
fmt.Println(reflect.TypeOf(&myFloat))
fmt.Println(reflect.TypeOf(&myBool))
```

输出结果：

```
*int
*float64
*bool
```



## 声明指针类型的变量（保存指针的变量）

由于指针是有类型的，因此可以声明指定指针类型的变量（声明保存指针的变量）。

```go
var myIntPointer *int
```

指针变量只能保存指向一种类型值的指针，因此变量可能只保存`*int`指针，只保存`*float64`指针，依此类推。

示例，声明保存指针的变量：

```go
var myInt int
fmt.Println(&myInt)                 //获取int类型的变量的地址（获取指针）
fmt.Println(reflect.TypeOf(&myInt)) //获取指向myInt的指针的类型（获取指针的类型）

var myIntPointer *int     //声明一个指向int的指针变量（声明*int类型的指针变量）
fmt.Println(myIntPointer) //未赋值的变量，值为nil

myIntPointer = &myInt //为指针变量分配一个指针（为指针变量赋值，值必须是一个指向相同类型的指针）
fmt.Println(myIntPointer)
```

上述代码中的myIntPointer就表示的是指针本身。

> 如果将一个指针直接赋值给了一个变量p，例如：
>
> ```go
> p := &x
> ```
>
> 它表示的意思是：p包含变量x的地址，或者p指向变量x（指针类型的零值是nil，测试p!=nil，结果为true，说明p指向一个变量（p的值是一个变量的地址））。



## 获取或更改指针引用的变量的值：*指针变量

注意：这里获取或更改的并不是指针变量本身，而是指针引用的变量对应的值，即指针指向的地址上的值。

例如：

```go
var myIntPointer *int  //声明一个指针，变量名叫myIntPointer
myIntPointer = &myInt  //该指针引用的变量为myInt
```

在上述代码中，myIntPointer是指针变量，而myInt是指针引用的变量。

在指针变量之前输入*运算符，来获得指针引用的变量的值。

因此，`*myIntPointer`，表示的是获取指针变量（myIntPointer）引用的变量（myInt）的值。==`*myIntPointer`读作“myIntPointer处的值”==。

示例一，获取myIntPointer处的值：

```go
var myInt int = 4
fmt.Println(&myInt)                 //获取int类型的变量的地址（获取指针）
fmt.Println(reflect.TypeOf(&myInt)) //获取指向myInt的指针的类型（获取指针的类型）

var myIntPointer *int    //声明一个指向int的指针变量（声明*int类型的指针变量）
myIntPointer = &myInt   //为指针变量分配一个指针的值（为指针变量赋值，值必须是一个指向相同类型的指针）
fmt.Println(myIntPointer)  //打印指针本身
fmt.Println(*myIntPointer) //打印指针处的值
```

myInt的值为4，指针变量myIntPointer引用的变量是myInt（见上述代码中的`myIntPointer = &myInt`），因此`*myIntPointer`的值也为4。输出结果如下：

```
0xc000014098
*int
0xc000014098
4
```

示例二，更新指针处的值：

```go
myFloat := 1.1
fmt.Println(myFloat)
myFloatPointer := &myFloat   //定义一个指针（值为变量myFloat的地址）
*myFloatPointer = 2.2        //给指针处的变量赋一个值（更新地址上的值)
fmt.Println(*myFloatPointer) //打印指针对应的地址上的值
fmt.Println(myFloat)         //指向该地址的变量的值也一并发生改变
```

输出结果：

```
1.1
2.2
2.2
```

由于myFloatPointer表示的是变量myFloat的地址，`*myFloatPointer=2.2`表示为该地址设置新值，而变量myFloat也指向该地址，因此一旦地址上的值发生了变化，对应的变量myFloat和`*myFloatPointer`的值都会变化。



> 在Go语言中，p指向的变量可以用`*p`来表示，也就是说这里的变量x可以直接通过`*p`来指代，换句话说，`*p`就代表着变量x，既然是变量，因此它也可以出现在赋值操作符左边，用于更新变量的值。
>
> ```go
> x := 1
> p := &x         //p是整型指针，指向x
> fmt.Println(*p) // 输出：1
> *p = 2          //*p指代的就是变量x，因此相当于是：x=2
> fmt.Println(x)  //输出：2
> ```
>
> 综上所述：
>
> `&x`：代表着指针本身，它的值是一个变量的地址。假如存在`p := &x`，那么p的值就是变量x的地址，此时，
>
> `*p`：代表着p指向的变量，即x变量本身，`*p`是变量x的别名，因此`*p`也被称为变量x的指针别名。



## 指针的比较运算

指针是可以比较的，两个指针**当且仅当**指向同一个变量或者两者都是nil的情况下才相等。

```go
var w, y int
fmt.Println(&w == &w, &w == &y, &y == nil) //输出：true false false
```

在Go语言中，函数返回局部变量的地址（指针）是非常安全的，例如：

```go
func wy() *int {
	v := 1
	return &v //返回一个指针
}
var p1 = wy()   //p1是一个指针
fmt.Println(p1) //输出：0xc00005c080
var p2 = wy()   //p2是一个指针
fmt.Println(p2) //输出：0xc00005c088
fmt.Println(p1 == p2) //输出：false
```

上述代码中，通过调用wy()产生的局部变量v即使在函数调用返回后依然存在，指针p1依然引用它，并且每次调用wy()都会返回一个不同的指针值。



## 函数指针

声明函数的返回类型是指针类型，并在函数的内部返回指针。

示例一，声明返回指针类型的函数：

```go
//声明一个返回类型为*float64的指针类型的函数
func createPointer() *float64 {
	var myFloat = 3.0
	return &myFloat //返回指针
}
```

调用部分：

```go
var myFloatPointer2 *float64 = createPointer()	//将返回的指针赋值给一个指针类型的变量
fmt.Println(*myFloatPointer2)  //输出3.0
```

在Go中，返回一个指向函数局部变量的指针是可以的。即使该变量不在作用域内，只要你仍然拥有指针，Go将确保你仍然可以访问该值。

示例二，将指针类型的变量作为参数传递给函数：

```go
func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer) //输出指针处的值
}
```

调用代码：

```go
var myBool bool = true
printPointer(&myBool) //向函数传递一个指针
```

可以将指针作为参数传递给函数，在函数的内部通过`*p`来获取指向的变量，从而能够实现让函数更新间接传递的变量值。

```go
//参数是一个整型指针
func incr(p *int) int {
	//获取指针指向的变量，并加1
	*p++
	return *p //返回指向的变量的值
}
m := 1
incr(&m)              //传入m指针，m加1
fmt.Println(m)        //输出：2
fmt.Println(incr(&m)) //再次传入m指针，m再次加1，并返回m的值，所以输出：3
fmt.Println(m)        //输出：3
```

上述代码中，`*p`是变量m的指针别名，指针别名允许我们不用变量的名字来访问变量。

下面是一个flag包使用指针相关的示例，自定义了两个标识参数，-n会忽略正常输出时结尾的换行符；-s使用sep替换默认参数输出时使用的空格分隔符。

```go
//echo4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)
//变量sep和n是指向标识变量的指针，因此必须通过*sep和*n来访问
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() //更新标识变量的默认值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("结束")
	}
}
```

执行结果：

```powershell
PS E:\go_cxsjyy\src\chapter2\echo4> go build
PS E:\go_cxsjyy\src\chapter2\echo4> .\echo4.exe a bc def
a bc def结束
PS E:\go_cxsjyy\src\chapter2\echo4> .\echo4.exe -s / a bc def
a/bc/def结束
PS E:\go_cxsjyy\src\chapter2\echo4> .\echo4.exe -n a bc def
a bc def
PS E:\go_cxsjyy\src\chapter2\echo4> .\echo4.exe -help
Usage of E:\go_cxsjyy\src\chapter2\echo4\echo4.exe:
  -n    omit trailing newline
  -s string
        separator (default " ")
PS E:\go_cxsjyy\src\chapter2\echo4>
```

总结：指针的作用就是实现其他语言（C#）中变量按照引用类型传递的方式。



## new函数

new函数用于创建指针。

```go
p := new(T)
```

表达式`new(T)`用于创建一个未命名的T类型的变量，并初始化为T类型的零值，并且返回其地址（指针），地址类型为`*T`。例如，当T为int时，`new(int)`返回的就是`*int`，表示的是整型指针。

```go
p := new(int)	//定义类型是整型指针（*int）的p，指针p指向未命名的int变量
fmt.Println(p)  //输出：0xc00005c080
fmt.Println(*p) //输出：0
*p = 2          //把未命名的int设置为2
fmt.Println(*p) //输出2
a := new(int)	//每一次调用new都返回一个具有唯一地址的不同变量
b := new(int)
fmt.Println(a) 	//输出：0xc00005c0b0
fmt.Println(b) 	//输出：0xc00005c0b8
```

使用new函数创建未知变量的指针时，由于不需要引入（和声明）一个虚拟的名字，因此可以直接在表达式中使用。

下面的两个`newInt`函数行为相同，相互等价：

```go
func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dumy int
	return &dumy
}
```

备注：new是一个预声明的函数，不是一个关键字。



## 总结

指针表示变量的地址。

- 声明指针类型的变量，用`*类型`，例如：`var myIntPointer *int`

- 表示指针的值，用地址运算符（`&变量名`）,例如：

  ```go
  var myInt int
  fmt.Println(&myInt)                 //获取int类型的变量的地址（获取指针）
  ```

- 为指针变量赋值，也用地址运算符（`&变量名`）,例如：`myIntPointer = &myInt`

- `&变量名`的形式，就代表着指针的值，表示指向该变量的地址。

- 获取或更改指针引用的变量的值，即更新指针指向的地址上的值，用`*指针变量`。

它们之间的关系，可以用如下代码说明：

```go
var myInt int			//声明普通变量
var myIntPointer *int	//声明*int类型（指向int的指针）的指针变量
myIntPointer = &myInt	//为指针变量赋值，此处也说明了&myInt表示的是指针的值
*myIntPointer=5			//设置指针引用的变量（myInt）的值，相当于为变量myInt赋值
```

或短变量声明指针：

```go
myFloat := 1.1
myFloatPointer := &myFloat   //定义一个指针（值为变量myFloat的地址）
*myFloatPointer = 2.2        //给指针处的变量赋一个值（更新地址上的值)
fmt.Println(*myFloatPointer) //打印指针对应的地址上的值
fmt.Println(myFloat)         //指向该地址的变量的值也一并发生改变
```



