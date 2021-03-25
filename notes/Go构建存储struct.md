# Go构建存储struct

类似于C#中的结构。

通常自定义的struct类型，都放在一个独立的包中，并且名称首字母大写（导出的），这样做，可以避免自定义类型名和其他变量名冲突。

## 声明 struct

方式一（不推荐）：使用struct关键字来声明一个struct类型。

```go
var myStruct struct {
	number float64
	word   string
	toggle bool
}
fmt.Printf("%#v", myStruct)
```

方式二（推荐）：通过type关键字定义一个struct类型。

使用变量的形式声明struct平时较少使用，常见的是通过type关键字，创建基于struct的类型。

### 使用type关键字将struct作为基础类型

为了定义一个类型，需要使用type关键字，后面跟着新类型的名字，然后是你希望基于的基础类型。

如果你使用struct类型作为你的基础类型，你需要使用struct关键字，后面跟着以花括号包裹的一组字段定义。

类型经常定义在函数外的包级别。

示例：

```go
//定义一个名为part的类型
type part struct {
	desc  string
	count int
}

//定义一个名为car的类型
type car struct {
	name     string
	topSpeed float64
}

func main() {
	var myCar car //定义一个car类型的变量
	myCar.name = "hi"
	myCar.topSpeed = 30.0
	fmt.Println(myCar)
}
```

这种方式声明的struct类型变量，就和c#中创建的类，类名即为struct类型的变量名。因此，上述的car就可以当作C# 中的一个类，可以通过car声明新的变量：

```go
var myCar car //定义一个car类型的变量
```

同时，使用方式也和C#中的类实例相似，可以将myCar作为实参，传入到形参是car类型的函数中。

```go
//定义一个名为car的类型
type car struct {
	name     string
	topSpeed float64
}
//定义一个形参为car类型的方法
func run(c car) {
	fmt.Println(c) //输出：{hi 30}
	c.name = "wt"
	c.topSpeed = 11.1
	fmt.Println(c) //输出：{wt 11.1}
}
func main() {
	var myCar car //定义一个car类型的变量
	myCar.name = "hi"
	myCar.topSpeed = 30.0
	fmt.Println(myCar) //输出：{hi 30}
	run(myCar)	//方法内部对变量值的修改并不会影响该变量的值，原因是go按照值类型传递
	fmt.Println(myCar) //仍然输出：{hi 30}
}
```

注意：Go是一个按值传递的语言，意味着函数调用时接收的是一个参数的拷贝。如果函数修改了参数值，它修改的只是拷贝，而不是原始值。

如要要按照引用类型传递值，需要借助指针来代替形参。

```go
//定义形参是指针类型的函数
func run(c *car) {
	c.name = "wt"
	c.topSpeed = 11.1
    (*c).name = "abc"
}
func main() {
	var myCar car //定义一个car类型的变量
	myCar.name = "hi"
	myCar.topSpeed = 30.0
	fmt.Println(myCar) //输出：{hi 30}
    //传入需要更新的值的指针
	run(&myCar)	//传入car struct类型的指针
	fmt.Println(myCar) //输出：{wt 11.1}
}
```

注意：使用点运算符在struct指针和struct上都可访问字段，因此在上述的run()函数中，可以直接通过`c.name = "wt"`来设置值，而不是`*c.name = "wt"`，如果非要使用“`*`”操作符，正确的写法应该是：`(*c).name = "abc"`。

虽然这种写法是指针的标准用法，但Go语言中的点运算符允许通过strcut的指针来访问字段，就像你可以通过struct值直接访问一样，可以不需要括号和*运算符。





## 获取和设置struct字段的值

```go
myStruct.number = 1     //设置struct字段值
myStruct.word = "hello" //获取struct字段值
fmt.Println(myStruct.number, myStruct.word)
```

