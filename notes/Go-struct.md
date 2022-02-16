# Go-struct

类似于C#中的结构或类。但Go中的struct是值类型。

通常自定义的struct类型，都放在一个独立的包中，并且名称首字母大写（导出的），这样做，可以避免自定义类型名和其他变量名冲突。



## struct 的声明与初始化

struct的声明有以下几种方式。

方式一（推荐），使用type关键字将struct作为基础类型：

```go
//定义一个名为car的类型
type car struct {
	name     string
	topSpeed float64
    //嵌套匿名结构体
    wy struct{
        name string
        age int
    }
}
```

方式二（不推荐）：

```go
var myStruct struct {
	number float64
	word   string
	toggle bool
}
fmt.Printf("%#v", myStruct)
//如果要同时打印字段名称，使用%+v
fmt.Printf("%+v", myStruct)
```

以上两种方式都是声明的是具名结构体。同时如方式一所示，可以在一个struct的字段中，定义匿名结构体字段。

struct的初始化操作有以下几种方式。

方式一：

```go
func main() {
	var myCar car //定义一个car类型的变量
	myCar.name = "hi"
	myCar.topSpeed = 30.0
	fmt.Println(myCar)
}
```

方式二，使用struct字面量形式初始化结构体：

```go
var myCat Cat = Cat{age: 10, name: "小白"}
myCat2 := Cat{age: 20, name: "小黑"}
```

方式三，使用new()函数通过获取指针的方式初始化结构体：

```go
var myCat3 *Cat = new(Cat)
(*myCat3).age = 30 // 可省略星号
myCat3.name = "小白" //省略了星号，go底层会对其进行处理，会加上(*myCat3)
fmt.Println(*myCat3)
```

如果使用new()方法来赋值：

```go
car2 := new(car)
```

相当于：

```go
car2 := &car{}
```

因为new()方法返回的是指针类型。

方式四，声明匿名struct并直接初始化：

```go
my := struct {
		name string
		age  int
	}{
		name: "张三",
		age:  10,
	}
	fmt.Println(my)
```

 

 

## 自定义类型

在Go语言中，既可以使用struct作为基础类型来定义类型，也可以基于int、string、bool或者其他任何类型来定义类型。

```go
//基于float64定义一个新类型
type Gallons float64
type Liters float64

func main() {
	var carFuel Gallons     //定义一个Gallons类型的变量
	carFuel = Gallons(10.0) //把float64转换为Gallons
	myLiter := Liters(11.1)
	fmt.Println(carFuel, myLiter)
}
```

可以把任何基础类型的值转换为定义的类型。像其他的类型转换一样，你写下需要转换到的类型，后面跟着在小括号中的你希望转换的值。

定义类型不能用来与不同类型的值一起运算，即使它们是来自相同的基础类型。

 



## 由struct组成的切片

```go
type wy struct {
	h int
	w int
}

a := []wy{
	{h: 3, w: 1},
	{h: 2, w: 3},
	{h: 4, w: 5}, //末尾必须有逗号
}
fmt.Println(a)
```



## 别名类型

别名类型与其源类型的区别恐怕只是在名称上，它们是完全相同的。

声明MyString是string类型的别名类型：

```go
type MyString = string
```

Go 语言内建的基本类型中就存在两个别名类型。byte是uint8的别名类型，而rune是int32的别名类型。

注意，如果是下述的这样声明：

```go
type MyString2 string  //注意：这里没有等号
```

MyString2和string就是两个不同的类型了。这里的MyString2是一个新的类型，不同于其他任何类型。

这种方式也可以被叫做对类型的再定义。对于这里的类型再定义来说，string可以被称为MyString2的潜在类型。潜在类型的含义是，某个类型在本质上是哪个类型。



## 匿名struct字段和嵌入struct

在定义struct类型时，允许定义匿名字段：struct字段没有名字，仅仅有类型。当声明一个匿名字段时，可以使用字段类型名称作为字段名称。

```go
type car struct {
	name     string
	topSpeed float64
}

type hello struct {
	car //匿名字段，其类型为car本身
	str string
}

func main() {
	//创建struct类型car的变量myCar，并同时为字段赋值
	myCar := car{name: "hi", topSpeed: 20}
	//为struct类型字段赋值
	hi := hello{car: myCar, str: "woqu"}
	fmt.Println(hi) //输出：{{hi 20} woqu}
}
```

> 使用匿名字段的方式将内部struct（这里的内部指的是上文中的car）增加到外部struct（这里的外部指的是上文中的hello），即内部struct作为了外部struct的匿名字段，被称为（struct）嵌入到了外部struct。你可以像访问外部字段一样访问嵌入的strcut字段。

一个struct作为另一个struct的匿名字段，称为嵌入struct，可以直接访问匿名字段对应的struct中的字段值。有点类似C#中的继承。

在上述代码中，hello包含了匿名字段car，可以直接使用hello的变量访问car的字段：

```go
//获取嵌入struct的字段名称
fmt.Println(hi.name)
fmt.Println(hi.topSpeed)
```

假如嵌入类型和父类型拥有相同的方法，那么父类型的方法的优先级高于嵌入类型的其他同名方法。



## struct类型的包导出

Go类型名称与变量和函数名称遵循相同的规则：如果变量、函数或者类型以大写字母开头，它就会被认为是导出的，并且可以从外部包来访问。

struct类型在其他包中也可以访问，需要满足如下条件：

- struct类型的名称首字母必须大写。
- struct类型的字段名称也必须首字母大写（没必要导出的字段首字母不用大写）。



## 方法

这里的方法，指的是自定义类型中定义的函数，和Python语言一样，类型中的函数被称为方法。类似于C#中的类的方法。

Go语言不支持方法重载。在同一个包中定义多个同名的函数不被允许，即使它们有不同类型的参数。你可以定义多个相同名字的方法，只要它们分别属于不同的类型。

方法和类型必须定义在同一包中。

方法名称以大写字母开头，则认为是导出的，如果它的名称以小写字母开头，则认为是不导出的。

就像其他的参数一样，接收器参数接收一个原值的拷贝。如果你的方法需要修改接收器，你应该在接收器参数上使用指针类型，并且修改指针指向的值。

Go语言没有为构造函数提供特殊的语言特性，构造函数和其他函数一样只是普通的函数。

### 自定义方法

假如存在如下自定义类型：

```go
//定义一个新的类型
type MyType string
```

需要为MyType添加自定义的方法，可以使用如下方式定义方法：

```go
//函数被定义在MyType上，m表示一个接收器
func (m MyType) sayHi(wy string) {
	fmt.Println(m)	//输出接收器参数的值
	fmt.Println(wy)
}
```

上述说明：

- m：m和MyType类似于函数参数的定义，m被称为接收器参数名称。这里表示MyType类型的接收器参数名称。

- MyType：接收器参数的类型。

上述方法的定义通常表达为：方法sayHi定义在MyType上。

一旦方法被定义在了某个类型，它就能被该类型的任何值调用。类似于C#中类的方法，可以被该类的任何实例对象调用一样。

方法调用：

```go
//创建一个MyType类型的值
value := MyType("Hello")
fmt.Println(value) 		//输出：Hello
value.sayHi("good") 	//调用MyType类型的syaHi方法
```

代码中的MyType("Hello")被称为方法接收器或接收者，接收器的值保存在变量value中，表示传递给接收器参数的接收器。

接收者简单来说就是`func (p *person) say()`中的`*person`，只不过它是一个指针类型的接收器。

### 接收器参数

接收器参数名称可以自己定义，通常使用接收器类型名称的首字母的小写形式作为名词。（上述中MyType的首字母小写是m），并且类型中定义的所有方法的接收器参数名称最好都保持一致，这样更易读。

接收器参数类似于C#中类的this对象，或其他语言中的“self”或者“this”。

> Go使用接收器参数来代替self和this。两者有着巨大的不同，self和this是隐含的，而你是显式地声明一个接收器参数。除此以外，接收器的用处相同，Go没有必要保留self和this关键字！（如果你想要，你可以将接收器参数命名为this，但是不要这么做，约定是使用接收器参数类型名称的第一个字母。）

### 指针类型的接收器参数

接收器参数与普通参数没有不同。但是就像其他任何参数，接收器参数接收一个接收器的拷贝值。如果你使用方法来修改接收器，你修改的是拷贝，而不是原始值。如果需要修改原始值，必须借助指针来完成。

例如：

```go
type Number int

//值类型传递
func (n Number) Double() {
	n *= 2
}

//指针类型传递值
func (n *Number) PointerDouble() {
	*n *= 2
}
```

调用方式：

```go
func main() {
	mynum := Number(4)
	pointer := &mynum     //指针
	mynum.PointerDouble() //值类型mynum自动转换为指针，调用指针类型方法
	fmt.Println(mynum)    //输出：8，同时指向该变量的指针对应的值都为8

	pointer.Double() //指针自动转换为值类型
	pointer.PointerDouble()
	fmt.Println(*pointer) //输出16
}
```

==当使用用一个非指针的变量调用一个需要指针的接收器的方法的时候，如上述代码中的`mynum.PointerDouble()`，Go会自动为你将非指针类型转换为指针类型。==这是因为Go语言在变量通过点标记调用方法的时候，会自动使用&取得变量的内存地址。所以就算不写`(&mynum).PointerDouble()`，代码也可以运行。

同样，当调用一个要求值类型的接收器，如上述代码中的`pointer.Double()`,指针类型也会自动转换为非指针类型，Go会自动帮你获取指针指向的值，然后传递给方法。

注意：只能获取保存在变量中的指针，也就是说，只有将接收器的值保存在变量中（如上述代码中的mynum），才能够进行指针转换。

如果直接使用下述方式，将会报错：

```go
Number(4).PointerDouble()
```

**使用指针作为接收者的策略应该是始终如一的。如果一种类型的某些方法需要用到指针作为接收者，就应该为这种类型的所有方法都使用指针作为接收者。**



### 方法与函数的区别

对于普通函数，接收者（这里指的是参数）是值类型时，不能将指针类型的数据直接传递，反之亦然。

对于方法，接收者（这里指类型接收器）是值类型时，可以直接用指针类型的变量调用方法，反过来也同样可以，例如：

```go
func (p Person) test3(){
}

p:=Person{}
// 此处使用指针类型的变量调用方法依然能够成功运行
(&p).test3() //此处仍然按值传递
```

但需要注意的是：无论外部是否使用指针类型变量调用方法test3()，最终实现按值传递还是引用传递的，都是依据接收器是否绑定的是指针类型来决定的。上述代码中，由于test3()的接收器使用的是Person，因此即使外面使用了指针变量进行调用，依旧按照值传递，除非将方法的声明改为：`func (p *Person) test3(){}`。具体见下文说明。



## 使用指针实现按引用传值

原理：==Go语言的函数和方法都以传值方式传递形参，这意味着函数总是基于被传递实参的副本进行操作。当指针被传递至函数时，函数将接收到传入内存地址的副本，在此之后，函数就可以通过解引用内存地址来修改指针指向的值。==

注意：Go是一个按值传递的语言，意味着函数调用时接收的是一个参数的拷贝。如果函数修改了参数值，它修改的只是拷贝，而不是原始值。

```go
func main() {
	type wy struct {
		h int
		w int
	}

	a := wy{1, 2}
	b := a
	b.h += 5
	fmt.Println(b) //输出：{6 2}
	fmt.Println(a) //输出：{1 2}
}
```

如要要按照引用类型传递值，需要借助指针来代替形参。

```go
//定义形参是指针类型的函数，这个地方直接影响是引用传递值还是值传递
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

由于在Go中，函数都是按照值传递的，即函数接收一个它们被调用时的参数的拷贝，即使它们是像struct那样的大值（一个有很多字段的大的struct），它会为原始的struct和被拷贝的struct都划分空间。

因此，除非struct只是一些小字段，==否则强烈建议向函数传入的是struct的指针==，而不是struct本身。当你传递一个struct指针的时候，内存中只有一个原始的struct，并且你可以读取它，修改它，或者做任何你想要的操作，都不会产生一个额外的拷贝。

综合示例：

```go
type car struct {
	name     string
	topSpeed float64
}

//方式一（不推荐）：直接返回struct类型，适用于小型struct
func createCar(name string, topSpeed float64) car {
	var oneCar car
	oneCar.name = name
	oneCar.topSpeed = topSpeed
	return oneCar
}

//方式二（推荐）：返回一个struct类型的指针
func createCar2(name string, topSpeed float64) *car {
	var oneCar car
	oneCar.name = name
	oneCar.topSpeed = topSpeed
	return &oneCar
}

//修改操作
func editCar(c *car, name string, topSpeed float64) {

	c.name = name
	c.topSpeed = topSpeed
}

//获取指针变量的值
func getCar(c *car) {
	fmt.Println("Name:", c.name)
	fmt.Println("TopSpeed:", (*c).topSpeed)
}

func main() {
	//方式一的调用
	myCar := createCar("wy", 11.1)
	fmt.Println(myCar) //输出：{wy 11.1}

	//方式二的调用
	myCar2 := createCar2("aa", 22.2)
	fmt.Println(myCar2) //输出：&{aa 22.2}

	//修改操作的调用
	editCar(myCar2, "bb", 33.3)
	fmt.Println(myCar2) //输出：&{bb 33.3}

	//获取指针变量的值
	getCar(&myCar)
	getCar(myCar2)
}
```



## struct 与 JSON ，结构标签（tag）

Go语言的json包要求结构中的字段必须以大写字母开头，并且包含多个单词的字段名称必须使用驼峰形命名惯例。如果需要让JSON数据使用其他格式形态，可以对结构中的字段打标签（tag），使json包在编码数据的时候能够按照我们的意愿修改字段的名称。

```go
func main() {
	type location struct {
		Lat, Long float64 //字段必须以大写字母开头
		Msg       string  `json:"message"`
	}

	curiosity := location{-4.5, 11.45, "你好"}

	//Marshal函数只对结构中被导出的字段实施编码
	bytes, err := json.Marshal(curiosity)
	if err != nil {
		os.Exit(1)
	}

	str := string(bytes)
	fmt.Println(str)
}
```

输出结果：

```
{"Lat":-4.5,"Long":11.45,"message":"你好"}
```

结构标签实际上就是一段与结构字段相关联的字符串。这里之所以使用被<code>``</code>包围的原始字符串字面量而不使用被<code>""</code>包围的普通字符串字面量，只是为了省下一些使用反斜杠转义引号的功夫而已。具体来说，如果我们把上例中的结构标签从原始字符串字面量改成普通字符串字面量，那么就需要把它改写成更难读也更麻烦的<code>"json:\"latitude\""</code>才行。

结构标签的格式为<code>key:"value"</code>，其中键的名称通常是某个包的名称。例如，为了定制<code>Lat</code>字段在JSON编码和XML编码时的输出，我们可以将该字段的结构标签设置成<code>`json:"latitude"xml:"latitude"`</code>。







## 组合

Go语言不支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承：

```go
import "image/color"

type Point struct {
	X,
	Y float64
}

type ColoredPoint struct {
	Point              // 嵌入匿名结构体
	Color color.RGBA   // 嵌入有名结构体
}

func main(){
    var myColor = &ColoredPoint{}
    //可以直接调用嵌入的结构体的字段
	myColor.X = 12.2
	myColor.Y = 12.4
	myColor.Color = color.RGBA{
		R: 120,
		G: 123,
		B: 100,
		A: 0,
	}
    
    	var myColor2 = ColoredPoint{
		Point{
			X: 0,
			Y: 0,
		},
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		},
	}
}
```

上述代码中，将Point嵌入ColoredPoint来提供X和Y这两个字段，这里将Point看作基类，把ColoredPoint看作Point的继承类或子类。

重点：

- 外部结构体可以使用内部嵌套的结构体的所有字段和方法，无论这些是否首字母大写都可以被调用。就像是嵌套的结构体完全属于本身一样。
- 当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。
- 结构体嵌入了两个（或多个）匿名结构体，如果两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错。
- 尽量不使用多重继承（嵌入多个匿名结构体）



## 结构体在内存中的结构

- 结构体的所有字段在内存中是连续的。
- 如果两个结构体要相互转换，需要满足条件：结构体中的字段名称、类型、个数都必须匹配。


