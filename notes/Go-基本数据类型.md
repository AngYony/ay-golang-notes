# Go - 基本数据类型

[TOC]

## 类型声明

type声明定义了一个新的命名类型，它和某个已有类型使用同样的底层类型。

命名类型的主要作用就是与底层类型的使用相互隔离，避免相互混乱使用。

即使两个命名类型使用了相同的底层类型，它们也不是相同的类型，仍然是不同的类型，所以它们不能使用算术表达式进行比较和合并。

命名类型的声明：

```
type name underlying-type
```

类型的声明通常出现在包级别，这里命名的类型在整个包中可见，如果名字是导出的（开头使用大写字母），其他的包也可以访问它。

命名类型的底层类型决定了它的结构和表达方式，以及它支持的内部操作集合，这些内部操作与直接使用底层类型的情况相同。

```go
package main

//定义两个类型
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	boilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func main() {
	//fmt.Println(Celsius(2.0) == Fahrenheit(2.0)) //编译错误
	fmt.Println(Celsius(2.0) == Celsius(Fahrenheit(2.0))) //输出true
    fmt.Println(float64(Celsius(2.0)) == float64(Fahrenheit(2.0))) //输出true
}
```

在上述代码中，即使Celsius和Fahrenheit都使用了相同的底层类型float64，它们也不是相同的类型，所以不能使用算术表达式进行比较和合并。要想解决这个问题，必须显式的进行类型转换。

对于每个类型T，都有一个对应的类型转换操作T(x)将值x转换为类型T。

如果两个类型具有相同的底层类型，或二者都是指向相同底层类型变量的未命名指针类型，则二者是可以相互转换的。

类型转换不改变类型值的表达方式，仅改变类型。

命令类型一般通常用于底层类型是复杂结构体类型，可以避免一遍遍地重复写复杂的类型。



## 查看数据类型

方式一：通过fmt.Printf()结合%T格式符来查看数据类型。

```
fmt.Printf("%T", i)
```

方式二：通过将任何值传递给reflect包的TypeOf函数，来查看它们的类型。

```go
package main
import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello ,Go!"))
}
```



## 基本类型

Go 的基本类型有：

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```



## 数值类型（number）

Go的数值类型包括：

- 整数
- 浮点数
- 复数

不同的数值类型分别有自己的大小。

### 整数

在Go语言中，整数类型分为：有符号整数和无符号整数。

| 内存占用情况  | 有符号整数 | 有符号取值范围         | 无符号整数 | 无符号取值范围 |
| ------------- | ---------- | ---------------------- | ---------- | -------------- |
| 8位（1字节）  | int8       | -128~127               | uint8      | 0~255          |
| 16位（2字节） | int16      | -32768~32767           | uint16     | 0~65535        |
| 32位（4字节） | int32      | -2147483648~2147483647 | uint32     | 0~4294967295   |
| 64位（8字节） | int64      | -2^63^~2^63^-1         | uint64     | 0~2^64^-1      |

除此之外，还有两种类型：int和uint。在特定平台上，其大小与原生的有符号整数/无符号整数相同，或等于该平台上的运算效率最高的值。int是目前使用最广泛的数值类型。int和uint的类型大小和平台相关（32位或者64位）。

rune类型是int32类型的同义词，通常用于指明一个值是Unicode码点（code point）。这两个名称可以互换。

byte类型是uint8类型的同义词，通常用于强调一个值是原始数据，而非量值。

uintptr：一种特殊的无符号整数，大小并不明确，但足以完整存放指针。uintptr类型仅仅用于底层编程。

int、uint 和 uintptr 都有别于其大小明确的相似类型的类型。（例如：int和int32是不同类型，尽管int天然的大小就是32位，并且int值若要当作int32使用，必须显式转换。）

**有符号整数的n位数字的取值范围是： -2^n-1^ ~ 2^(n-1）^-1，由于需要保留最高位作为符号位来表示正数或者复数，因此在次方上减1。例如，int8的取值范围为：-128 ~ 127**。

**无符号整数的n位数字的取值范围是： 0 ~ 2^n^-1，无符号整数都是非负数。例如，uint8的取值范围为：0 ~ 255**。

#### 二元运算符

Go的二元操作符包括算术、逻辑、比较、位运算等，按优先级的降序排列如下：

1.   `*` 	`/`	`%`	`<<`	`>>`	`&`	`&^`
2.   `+`	 `-`    `|`    `^`
3.   `==`   `!=`  `<`    `<=`    `>`    `>=`
4.   `&&`
5.   `||`

二元运算符分五大优先级，同级别的运算符满足左结合律，可以通过圆括号按指定次序计算。

在Go语言中，取模运算符 % 仅能用于整数，用于取余计算。并且，**取模余数的正负号总是与被除数一致**，例如：-5%3 和 -5%-3 计算结果都是 -2。整数的除法运算（/）结果仍然是整数（会舍弃小数部分），例如 5/4结果为1。

在Go语言中，整数类型发生了溢出时（算术运算结果所需的位超出该类型的范围），溢出的高位部分会无提示地丢弃，并不会报异常。假如原本的计算结果是有符号类型，且最左侧位是1，则会形成负值，例如：

```go
var i int8 = 127
fmt.Println(i + 1) //结果为：-128
fmt.Println(i * i) //结果为：1

var u uint8 = 255
fmt.Println(u + 1) //结果为：0
fmt.Println(u * u) //结果为：1

var u2 uint8 = 0
fmt.Println(u2 - 1) //结果为：255
```



#### 无符号整数（uint?）注意事项

- 整数类型（包括无符号整数）在算术溢出的时候，可能会得到意想不到的值，因此无符号整数往往**只用于**位运算符和特定算术运算符。（关于位运算符参见笔记《二进制、八进制、十六进制、位运算符》）
- 即使要表示非负数，也最好不要使用无符号整数，原因是无符号整数一旦牵扯到算术运算，往往会得到意想不到的结果。例如无符号整数值为0，减一之后，将变成对应位的最大值，如上述示例中的u2。
- 无符号整数最好只用于位运算符，除此之前，尽量不要使用。

#### 其他类型转换为整型

浮点型转成整型，会舍弃小数部分，趋零截尾（正值向下取整，负值向上取整）。



### 浮点数

- 浮点数存放形式：浮点数=符号位+指数位+尾数位
- 尾数部分可能丢失，造成精度损失

Go具有两种大小的浮点数：float32和float64。

在Go语言中，所有带小数点的数字在默认情况下都会被设置为float64类型。

除非有特殊理由，否则应该优先使用float64类型。

常量math.MaxFloat32是float32的最大值；常量math.MaxFloat64是float64的最大值。

```go
var f1 float32 = 999999
fmt.Println(f1) //输出：999999

var f2 float32 = 999999 + 1
fmt.Println(f2) //输出：1e+06

var f3 float32 = 16777216 //临界值
fmt.Println(f3 == f3+1)   //输出：true
```

十进制下，float32的有效数字大约是6位，float64的有效数字大约是15位。（如上述代码中的f1和f2）。

注意：float32能精确表示的正整数范围有限（大约<=16777216），因此应优先使用float64。

对于float64类型，最小值为2^-52^，而对于float32类型，最小值为2^-23^。

小数点前的数字可以省略（.707），后面的也可以省去（1.）。非常小或非常大的数字最好使用科学计数法表示。

5.1234E2表示：5.1234乘以10的2次方，也就是512.34。

5.1234E-2表示：5.1234除以10的2次方，也就是0.051234。

```go
fmt.Println(5.1234E2)  //输出：512.34
fmt.Println(5.1234e-2) //输出：0.051234
```



用于输出浮点值的Printf谓词有：

- %g：该谓词会自动保持足够的精度，并选择最简洁的表示方式。
- %e：按照有指数的形式表示数据表。
- %f：按照无指数的形式表示数据表。

下述代码按8个字符的宽度输出自然对数e的各个幂方，结果保留三位小数：

```go
for x := 0; x < 8; x++ {
	//fmt.Println(math.Exp(float64(x))) //输出自然对数e的各个幂方值
	fmt.Printf("x=%d e^x=%8.3f \n", x, math.Exp(float64(x)))
}
```

输出结果：

```
x=0 e^x=   1.000 
x=1 e^x=   2.718 
x=2 e^x=   7.389 
x=3 e^x=  20.086 
x=4 e^x=  54.598 
x=5 e^x= 148.413 
x=6 e^x= 403.429 
x=7 e^x=1096.633 
```



注意：不能直接将两个浮点数进行比较，比较结果永远是false。

```go
a := 0.1
a += 0.2
fmt.Println(a)
fmt.Println(a == 0.3)
```

输出结果：

```
0.30000000000000004
false
```

推荐的做法是，计算它们之间的差值，然后通过判断这个差值的绝对值是否足够小来判断两个浮点数是否相等。

```go
fmt.Println(math.Abs(a-0.3) < 0.00001)
```



#### math包中的特殊值

正无穷大（+Inf）和负无穷大（-Inf）：表示超出最大许可值的数及除以零的商。

NaN（Not a Number）：表示数学上无意义的运算结果（如0/0，或Sqrt(-1)）。

math.IsNaN 函数判断其参数是否是非数值。

math.NaN 函数则返回非数值（NaN）。

在数字运算中，我们倾向于将NaN当作信号值，但直接判断具体的计算结果是否为NaN可能导致潜在错误，因为与NaN的比较总不成立。

### 复数

Go具备两种大小的复数：complex64和complex128，二者分别由float32（实部和虚部）和float64（实部和虚部）构成。

通过内置的complex函数，根据给定的实部和虚部创建复数。

提取复数的实部时，使用内置的real函数；提取复数的虚部时，使用内置的imag函数。

如果在浮点数或十进制数后面紧接着写字母i，如3.12i或2i，它就变成了一个虚数。

```go
func main() {
	c := 3 + 4i
	fmt.Printf("%v  %[1]T \n", c) //输出：(3+4i)  complex128
	//计算绝对值
	fmt.Println(cmplx.Abs(c)) //输出：5
	//直接使用i会提示变量未定义，因此使用1i代替i。
	var c2 = cmplx.Pow(math.E, 1i*math.Pi) + 1
	//也可以使用下述语句来简化写法
	c3 := cmplx.Exp(1i * math.Pi) //表示E的多少次方
	fmt.Println(c2, c3)           //输出的结果相同
}
```





### big包

当一个数字，超过int64能够存储的上限时，需要借助big包进行数据存储。

big包提供了以下3种类型：

- 存储大整数的big.Int
- 存储任意精度浮点数的big.Float
- 存储诸如1/3的分数的big.Rat。

一旦决定使用big包中的类型，就需要在等式的每个部分都使用这种类型，即使对已存在的常量来说也是如此。

创建一个big.Int类型的变量并赋值：

```go
lightSpeed := big.NewInt(299792)
```

NewInt方法只能传入int相关的类型的值，一旦是一个超出int范围的值，就需要借助其他方法来创建：

```go
distance2 := new(big.Int)
//由于传入的数值是10进制的，所以第二个参数指定为10
distance2.SetString("240000000000000000000000000000000000000", 10)
```

综合示例：

```go
var distance int64 = 41.3e12 //表示41.3乘以10的12次方
fmt.Println(distance)
lightSpeed := big.NewInt(299792)
secondsPerDay := big.NewInt(86400)
distance2 := new(big.Int)
//由于传入的数值是10进制的，所以第二个参数指定为10
distance2.SetString("240000000000000000000000000000000000000", 10)
fmt.Println(distance2)
seconds := new(big.Int)
//进行除法运算，并赋值给seconds
seconds.Div(distance2, lightSpeed)
days := new(big.Int)
days.Div(seconds, secondsPerDay)
fmt.Println(days)
```





## 布尔值

在Go语言中，&&的优先级比||高。



## 字符串和byte类型

字符串和byte类型的相关介绍见《Go字符串》。

```go
// 当直接输出byte时，相当于输出的是其对应的字符的ASCII码值
var b byte = 'A'
fmt.Printf("%T \n", b) // 输出：uint8
fmt.Println(b)         // 输出：65
fmt.Printf("%T \n", 'A') // 输出：int32
fmt.Println('A')         // 输出：65

// 使用格式化字符原样输出
fmt.Printf("%c \n", 'A') // 输出：A

// 对于byte不能存储的字符
fmt.Printf("%c %T %d \n", '王', '王', '王') // 输出：王 int32 29579 
var y uint32 = '你'
fmt.Printf("%c %T %d \n", y, y, y) // 输出：你 uint32 20320

// 报错
var x byte = '我' //编译无法通过
fmt.Printf("%c \n", x) // 输出：A
```

byte = uint8 = 1字节，表示的范围：0~255

对于存储超过一个字节的字符，需要使用其他类型，如：int16/uint16（2byte）、int32/uint32（4byte）、int64/uint64（8byte）。

正因为如此，存储汉字时，需要3个字节，因此可以采用int32/uint32来进行存储，在go语言中，约定使用rune类型存储字符，rune是int32的别名，也是基本数据类型。

字符类型可以进行计算：

```go
var x byte = 'A'
fmt.Println(x + 10)     // 输出：75
fmt.Println('b' + 3)    // 输出：101
fmt.Printf("%c", 'b'+3) // 输出：e
```

字符串存储的本质：

字符 -> 对应码值 -> 二进制 -> 存储











