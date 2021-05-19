# Go - 封装

将程序中的数据隐藏在一部分代码中而对另一部分不可见的方法称为封装。在Go中使用未导出的变量、struct字段、函数或者方法，把数据封装在包中。

这里的封装和C#中类的封装类似，包括：

- 设为私有变量（字段名非大写）
- 添加setter和getter方法
- 将类型放入到另一个包中，并将数据字段设置为非导出的。



Go开发者通常在需要的时候才使用封装，比如字段数据需要被setter方法校验时。在Go中，如果你不需要封装一个字段，通常导出并且允许直接访问它。



## 封装操作

### 可见性

每一个包给它的声明提供独立的命名空间。可以通过控制变量在包外面的可见性或导出情况来隐藏信息：以大写字母开头的标识符对其他包可见。

在下面的示例中，两个不同的.go文件都在同一个包中，temps.go文件中的代码如下：

```go
//定一个包和包级别的常量和类型声明，并且首字母均大写
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g℉", f)
}
```

由于上述代码定义的是包级别的常量和类型声明，因此它们可以在同一个包中直接被引用，并且它们的名字都是以大写字母开头的，因此也可以在其他包中被引用。

conv.go文件中的代码如下：

```go
//定义温度相互转换的方法
package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
```

如果不是在同一个包中调用，而是在其他包中使用上述定义的常量和类型声明，需要导入其所在的包，并且通过包名的形式进行访问。

例如，下述代码在main包中调用tempconv包中的成员，此时需要使用`tempconv.成员`的形式进行访问：

```go
//测试多个.go文件在同一个包中的调用情况
package main

import (
	"chapter2/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("TTTT! %v\n", tempconv.AbsoluteZeroC) //输出：TTTT! -273.15℃
	fmt.Println(tempconv.CToF(tempconv.BoilingC))    //输出：212℉
}
```

package声明的前面通常需要写明对整个包进行描述的文档注释。并且每一个包里只有一个文件应该包含该包的文档注释。扩展的文档注释通常放在一个文件中，按惯例名字叫做doc.go。

注意：

- 包级别的成员，在其所在的包中的任意文件都可见，可以直接通过名称来引用。
- 如果包级别的声明以大写字母开头，那么它将在其他包中也可见，可以在其他包中通过`包名.成员`的形式进行引用。

### 定义类型

将类型放入到另一个包中，并将数据字段设置为非导出的。只需在字段声明和它出现的所有位置将首字母修改为小写字母即可。

```go
package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}
```

### 定义setter

依照惯例，Go的setter方法名为SetX的形式。

setter方法是用来设置字段或者基础类型中的其他值的方法。

setter方法需要指针接收器。

未导出的变量、struct字段、函数、方法等仍然能够被相同包的导出的函数或者方法访问。

setter示例：

```go
func (d *Date) SetYear(year int) error {
	if year < 1 {
		//如果year的值是否非法的，返回一个错误
		return errors.New("年份错误")
	}
	d.year = year //d.Year会自动取到指针指向的值（就像我们使用了（*d）.Year）。
	return nil 	  //返回nil作为错误
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("月份错误")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day值错误")
	}
	d.day = day
	return nil
}

```

### 定义getter

按照惯例，getter方法的名称应该与访问的字段或者变量的名字相同。

当然，如果你希望方法被导出，它的名字的首字母需要大写。

通常，如果类型的任何方法接受接收器指针类型，为了一致性，通常来说所有的方法都应该这样做。由于我们必须对所有的setter方法使用接收器指针，我们也应对所有的getter方法使用指针。

注意：Go社区已经在一个大会上决定了在getter方法前面去掉Get前缀，所以不要将getter方法的名称以“Get”开头。确保遵守约定：一个getter方法应该与它访问的字段名称相同，如果函数需要导出，它的首字母需要大写）

getter示例：

```go
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

```



### 调用

```go
package main

import (
	"book_head_first_go/ch10/calendar" //导入新包
	"fmt"
)

func main() {

	date := calendar.Date{}
	err := date.SetYear(2018)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err) //报告错误，并停止程序运行
	}

	err = date.SetMonth(10)
	if err != nil {
		fmt.Println(err)
	}

	date.SetDay(13)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)
}
```

## 封装中的嵌入

一个类型使用匿名字段的方式保存到另一个struct类型中，被称为嵌入了struct。

嵌入类型的导出方法会提升到外部类型。它们可以被调用，就像它们是在外部类型上定义的一样。

外部类型定义的方法和内部嵌入类型的方法的生存时间是一样的。

一个嵌入类型的未导出方法不会被提升到外部类型。

关于嵌入的介绍，见struct篇。



