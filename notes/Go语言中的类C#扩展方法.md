# Go语言中的类C#扩展方法

### Go语言中的类C#扩展方法

在C#语言中，可以定义类型的扩展方法，在Go语言中，也提供了这种方式。例如：

```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
```

上述代码中，Celsius类型的参数c出现在了函数String()名字的前面，表示名字叫String的方法关联到Celsius类型，类似于C#中类型的扩展方法，这样每次调用Celsius类型的变量时，就可以调用它的扩展方法String()了。

```go
c := FToC(212.0)		//先转换为Celsius类型
fmt.Println(c)          //输出100℃，Println()需要传入一个字符串，隐式被调用
fmt.Println(c.String()) //输出100℃
fmt.Printf("%v\n", c)   //输出100℃，不需要显式调用字符串
fmt.Printf("%s\n", c)   //输出100℃
fmt.Printf("%g\n", c)   //输出100，不调用字符串
fmt.Println(float64(c)) //输出100，不调用字符串
```

