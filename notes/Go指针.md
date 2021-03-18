# Go指针

指针：表示变量地址的值称为指针，指针指向变量的位置（地址）。



## 指针的值的表示方式：地址运算符（&变量名）

地址运算符：`&变量名`（可以使用一个`&`符号获取变量的地址）。

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

综合示例：

```

```





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

