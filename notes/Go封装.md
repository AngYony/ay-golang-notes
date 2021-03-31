# Go封装

将程序中的数据隐藏在一部分代码中而对另一部分不可见的方法称为封装。在Go中使用未导出的变量、struct字段、函数或者方法，把数据封装在包中。

这里的封装和C#中类的封装类似，包括：

- 设为私有变量（字段名非大写）
- 添加setter和getter方法
- 将类型放入到另一个包中，并将数据字段设置为非导出的。



Go开发者通常在需要的时候才使用封装，比如字段数据需要被setter方法校验时。在Go中，如果你不需要封装一个字段，通常导出并且允许直接访问它。



## 封装操作

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



