# Go - 组合

Go语言不支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承：

```go
import "image/color"

type Point struct {
	X,
	Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
```

上述代码中，将Point嵌入ColoredPoint来提供X和Y这两个字段，这里将Point看作基类，把ColoredPoint看作Point的继承类或子类。

