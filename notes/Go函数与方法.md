# Go函数与方法

## 函数

### 可变长参数函数

一个可变长参数函数可以以多种参数个数来调用。为了让函数的参数可变长，在函数声明中的最后的（或者仅有的）参数类型前使用省略号（...）。

可变长参数函数的最后一个参数接收一个切片类型的变长参数，这个切片可以被函数当作普通切片来处理。

仅仅函数定义中的最后一个参数可以是可变长参数；你不能把它放到必需参数之前。

```go
func myFunc(a int, b ...string) {
	fmt.Println(a, b)
}
```

调用时，只需要为可变参数传入0个或多个值即可：

```
myFunc(1, "a", "b")
```

如果在调用可变长参数函数时，想要为可变参数传入一个切片变量，需要在传入的切片变量后增加省略号（...)。例如：

```go
wy := []string{"AA", "BB"}
myFunc(1, wy...)   //使用切片变量代替多个参数值的调用
```



## 方法

这里的方法，指的是自定义类型中定义的函数，和Python语言一样，类型中的函数被称为方法。类似于C#中的类的方法。

关于方法的定义和使用，见struct笔记部分。

Go语言不支持方法重载。



## 函数作为类型（一级函数）

在具有一级函数的语言中，可以将函数赋值给变量，然后使用这些变量来调用函数。

类似于JavaScript中的，将函数分配给一个变量，然后从这些变量调用函数。也可以将函数作为另一个函数的参数进行传递。类似于C#中的委托。

```go
func syaHi() {
	fmt.Println("hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}
func main() {
	var myFunc func()
	myFunc = syaHi
	myFunc()

	twice(sayBye)
}
```

函数的参数和返回值是其类型的一部分。

保存函数的变量需要指定函数应该具有哪些参数和返回值。该变量只能保存参数的数量和类型以及返回值与指定类型匹配的函数。

```go
func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}
func main() {
	var greeterFunction func()
    //定义个函数类型的变量，类似于C#中的委托
	var mathFunction func(int, int) float64

	greeterFunction = sayHi
	mathFunction = divide

	greeterFunction()
	fmt.Println(mathFunction(4, 2))
}
```

当调用其他函数时，函数也可以作为参数传递。

上述代码，可以将matchFunction变量作为接受函数的形参使用，这样就可以实现限定传入函数应该具有的参数和返回类型。

```go
func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

//定义一个指定传入的函数类型
func doMath(mathFunction func(int, int) float64) {
	result := mathFunction(10, 2)
	fmt.Println(result)
}
func main() {
	var greeterFunction func()
	greeterFunction = sayHi
	greeterFunction()

	//定义函数类型变量
	var mathFunction func(int, int) float64
	mathFunction = divide //为变量分配函数
	//将函数类型变量作为参数传递给以函数作为形参的函数，类似于C#委托
	doMath(mathFunction)
}
```

