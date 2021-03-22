# Go包文件

在Go语言中，包的作用类似C#中的命名空间，一个包的源代码保存在一个或多个以.go结尾的文件中，文件所在的完整目录名的尾部（$GOPATH/src/之后的）就是包导入的路径。

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

### 导入

在Go程序里，每一个包通过称为`导入路径`的唯一字符串来标识。

```go
package main

import (
	"chapter2/tempconv"
	"fmt"
)
```

一个导入路径标注一个目录，目录中包含构成包的一个或多个Go源文件。

除了导入路径之外，每个包还有一个包名，它以短名字的形式（且不必是唯一的）出现在包的声明中。

按约定，包名匹配导入路径的最后一段，例如“chapter2/tempconv”的包名是tempconv。这种由导入声明给导入的包绑定的短名字，可以用来在整个文件中引用包的内容。为了避免冲突，导入声明还可以设定一个可选的名字。

如果导入一个没有被引用的包，将会触发一个编译错误。这个检查有助于消除代码演进过程中不再需要的依赖。

### 包的初始化

包的初始化按照如下过程进行：

1. 首先初始化包级别的变量，这些变量按照声明顺序初始化，在依赖已解析完毕的情况下，根据依赖的顺序进行。例如：

   ```go
   var a = b + c //3：最后把a初始化为3
   var b = f()   //2：接着通过调用f()将b初始化为2
   var c = 1     //1：首先初始化c为1
   func f() int  { return c + 1 }
   ```

2. 如果包由多个.go文件组成，初始化按照编译器收到文件的顺序进行：go工具会在调用编译器前将.go文件进行排序。

   对于包级别的每一个变量，生命周期从其值被初始化开始，但是对于其他一些变量，比如数据表，初始化表达式不是简单地设置它的初始化值，而是需要调用init()函数（见下述）。

3. 包的初始化按照在程序中导入的顺序来进行，依赖顺序优先，每次初始化一个包。例如，如果包p导入了包q，可以确保q在p之前已完全初始化。初始化过程是自下向上的，main包最后初始化。在这种方式下，在程序的main函数开始执行前，所有的包已初始化完毕。

#### init()

任何文件都可以包含任意数量的init函数，它的声明形式如下：

```go
func init(){...}
```

init函数不能被手动的调用和被引用，另一方面，它也是普通的函数。在每一个文件里，当程序启动的时候，init函数按照它们**声明的顺序**自动执行。

