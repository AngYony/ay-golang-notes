# Go - 函数、方法与闭包



## 函数

在Go语言中，使用关键字func定义函数。

Go语言中的函数注意事项：

- 不支持同名函数重载
- 不支持默认参数
- 属于第一类对象，即：函数可以在运行期创建，可以作为其他函数的参数或返回值，可以存入变量的实体中，如匿名函数。
- 具有相同的参数和返回值列表的函数，属于同一类型。（var wy=**func (a int) int**{}）
- 函数类型的变量，只能和nil进行比较操作，用来判断是否为nil，而不能和其他任何类型进行比较运算。
- 函数的参数都是值拷贝传递，当拷贝的是目标对象时，就是值传递，当拷贝的是指针时，就是引用传递。
- 支持可变参数，且可变参数必须放在参数列表的最后。





### 函数声明

在Go语言中，函数是第一类对象，可以将函数保存到变量中。

函数主要有具名和匿名之分，包级函数一般都是具名函数，具名函数是匿名函数的一种特例。

Go语言中每个类型还可以有自己的方法，方法其实也是函数的一种。

#### 具名函数

```go
func Add(a, b int) int{
	return a+b
}
```

#### 匿名函数

匿名函数也就是没有名字的函数，在Go中也被称为函数字面量。匿名函数通常赋值给某个变量，匿名函数需要保留对该变量（位于匿名函数外部作用域中的变量）的引用，以便通过该变量调用函数，所以匿名函数都是闭包的。关于闭包的详细介绍，见下文。

```go
var Add = func(a, b int) int{
	return a+b
}
r:=Add(1,2)
```

匿名函数可以在声明的时候就直接运行：

```go
func(n1 int) {
	fmt.Println(n1)
}(3)
```



#### 多个参数和多个返回值

Go语言中的函数可以有多个参数和多个返回值，==参数和返回值都是以**传值的方式**和被调用者交换数据==。

```go
func Swap(a, b int)(int, int) {
	return b, a
}
```

#### 可变长参数函数

在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。

为了让函数的参数可变长，在函数声明中的最后的（或者仅有的）参数类型前使用省略号（...）。

可变长参数函数的最后一个参数接收一个切片类型的变长参数，这个切片可以被函数当作普通切片来处理。

仅仅函数定义中的最后一个参数可以是可变长参数；你不能把它放到必需参数之前。

```go
//b对应[]int切片类型
func myFunc(a int, b ...int) {
	fmt.Println(a, b)
}
```

调用时，只需要为可变参数传入0个或多个值即可。

如果在调用可变长参数函数时，想要为可变参数传入一个切片变量，需要在传入的切片变量后增加省略号（...)。例如：

```go
wy := []string{"AA", "BB"}
myFunc(1, wy...)   //使用切片变量代替多个参数值的调用
```

当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：

```go
func wyPrint(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	var a = []interface{}{123, "abc"}
    //传入的参数是a...等价于直接调用Print(123,"abc")
	wyPrint(a...) //输出：123 abc
    //传入的是未解包的a，等价于直接调用Print([]interface{}{123, "abc"} )
	wyPrint(a)	  //输出：[123 abc]
}
```

#### 给返回值命名的函数

```go
//在map中查询指定key的value
func find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}
```

如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值：

```go
func inc() (v int) {
	defer func() { v++ }()
	return 42
}
```



### 一等函数

在Go语言里面，函数是一等值，它可以用在整数、字符串或其他类型能够应用的所有地方：你可以将函数赋值给变量，可以将函数传递给函数，甚至可以编写创建并返回函数的函数。

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



### 声明函数类型

声明函数类型指的是以某个函数作为底层类型来声明新的类型。

例如，以`func() int`作为底层类型，声明新的函数类型wyfunc：

```go
type wyfunc func() int
```

这样，当存在如下形参的函数时：

```go
func say(count int, wy func() int)
```

就可以改写为：

```go
func say(count int, wy wyfunc)
```



## 函数调用机制

当调用的是值传递参数的函数时：

1. 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间和其他的栈的空间区分开来；
2. 在每个函数对应的栈中，数据空间是独立的，不会混淆；
3. 当一个无返回值的函数调用完毕（执行完毕）后，程序会销毁这个函数对应的栈空间。









## 方法

这里的方法，指的是自定义类型中定义的函数，和Python语言一样，类型中的函数被称为方法。类似于C#中的类的方法。

关于方法的定义和使用，见struct笔记部分。

Go语言不支持方法重载。







## 闭包

闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）。

闭包通常发生在匿名函数。

```go
func main() {
    //定义一个匿名函数
	f := func(fname string) {
		fmt.Println("匿名函数：", fname)
	}
	f("函数A")
}
```

匿名函数赋值给了位于外部作用域中的变量f，当需要执行匿名函数时，匿名函数必须保留对变量f的引用，因此匿名函数都是闭包的。

闭包示例：

```go
//定义一个新的函数类型sayF
type sayF func() string

//形式是sayF类型的变量，并且返回的也是sayF类型函数
func hello(s sayF, riyu string) sayF {
	return func() string {
		return s() + riyu
	}
}

func english() string {
	return "英语"
}

func main() {
	//获取返回的函数赋值给变量
	lb := hello(english, "日语")
	//执行变量对应的函数体
	fmt.Println(lb())
}
```

上述代码中，lb引用了被hello函数用作形参的s变量和riyu变量，尽管hello函数已经返回了，但是被闭包捕获的变量将继续存在，因此调用lb仍然能够访问这两个变量。

术语闭包就是由于匿名函数封闭并包围作用域中的变量而得名的。

注意：闭包保留的是周围变量的引用而不是副本值。闭包对捕获的外部变量并不是以传值方式访问，而是以引用方式访问。

闭包的示例二：

```go
// 定义一个返回值类型是函数的函数
func addNum() func(int) int {
	var n int = 100
    // 返回的是一个匿名函数，但在函数内部引用了一个外部变量n，匿名函数和n形成了一个整体，构成了闭包
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	f := addNum()
	fmt.Println(f(1)) // 输出：101
	fmt.Println(f(2)) // 输出：103
	fmt.Println(f(3)) // 输出：106
}
```

闭包有点类似面向对象中的类，函数是类的操作方法，而上述的变量n相当于类的字段。当反复调用f函数时，因为n只初始化的一次，闭包保留的是周围变量的引用而不是副本值，因此每次调用一次就进行累加，而不是重新初始化。

要搞清楚闭包的关键，就是分析出返回的函数使用到了哪些周边变量，因为函数和变量共同组成了一个整体构成了闭包。



为防止闭包引发问题，最好的做法时候，对函数进行参数传递。

规律：

只要出现func，就一定要考虑闭包问题，一旦函数体内引入了外部的变量，就必须要通过函数形参的方式来传递值。



