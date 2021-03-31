# Go-panic-recover

不建议使用panic和recover。Go开发人员被鼓励以处理程序其他部分的方式处理错误：使用if和return语句，以及error值。

> 不要在Go中寻找等同于其他编程语言中的exception的东西。这个特性被故意省略了。

## panic

panic：宕机，指发生了能够使程序崩溃的异常。

panic函数需要一个满足空接口的参数（也就是说，它可以是任何类型）。该参数将被转换为字符串（如果需要），并作为panic日志信息的一部分打印出来。

可以通过简单地调用内置的panic函数来引发panic：

```go
func main() {
	fmt.Println("1")
	panic("完犊子")
	fmt.Println("2")
}
```

Go保持一个调用堆栈，即在任何给定点上处于活动状态的函数调用的列表。当程序发生panic时，panic输出中包含堆栈跟踪，即调用堆栈列表。

panic通常用于不可能的情况，大多数情况下，仍然使用error。

### defer和panic

当程序出现panic时，所有延迟的函数调用仍然会被执行。如果有多个延迟调用，它们的执行顺序将与被延迟的顺序相反。

```go
func one() {
	//这个函数调用首先被延迟，所以它将在最后执行。
	defer fmt.Println("defer-one")
	two()
}

func two() {
	//这个函数调用最后被延迟，所以它将首先执行
	defer fmt.Println("defer-two")
	panic("two引发了崩溃")
}

func main() {
	one()
}
```

输出：

```
defer-two
defer-one
panic: two引发了崩溃

goroutine 1 [running]:
main.two()
...
```



## recover

panic会导致程序崩溃，并出现难看的堆栈跟踪信息，如果只想展示错误信息，可以使用recover函数。

在正常程序执行过程中调用recover时，它只返回nil，而不执行其他操作。

```go
fmt.Println(recover())	//输出nil
```

recover函数返回最初传递给panic函数的任何值。

一旦在调用panic函数之后，程序将处于panic状态，panic之后的所有代码都不能执行，所以recover方法只能出现在panic语句之前。因此，必须借助defer来实现延迟调用，从而实现在程序陷入panic时，仍然可以调用recover方法。

不能直接对recover函数使用defer关键字，可以在一个单独的函数中放置一个recover调用，并在引发panic的代码之前使用defer调用该函数。

```go
func calmDown() {
    //recover返回传递给panic的任何值
	fmt.Println(recover())	//调用recover，并打印panic值
}

func freakOut() {
	defer calmDown() //延迟对函数的调用
	panic("崩溃了。。。")
	fmt.Println("panic之后的代码永远不会被执行")
}

func main() {
	freakOut()
	//这段代码在freakOut2返回之后会运行
	fmt.Println("运行完成")
}
```

freakOut函数内部调用了recover()和panic，panic之后的代码不会被执行，但是，在产生panic的freakOut函数返回之后，正常的执行将恢复，因此最后一个语句能够正常执行。

上述整个代码输出结果：

```
崩溃了。。。
运行完成
```

当没有panic时，调用recover返回nil。当出现panic时，recover返回传递给panic的任何值。

recover的返回值的类型也是interface{}。你可以将recover的返回值传递给诸如Println（它接受interface{}值）之类的fmt函数，但是你不能直接对其（recover的返回值）调用方法，必须使用类型断言将其转换回其底层类型。

```go
func calmDown2() {
    p := recover()			//返回一个interface{}值
	err, ok := p.(error)	//断言panic值的类型为error
	if ok {
		fmt.Println(err.Error())	//调用error的Error方法
	}
}

func main() {
	defer calmDown2()
	err := fmt.Errorf("错误消息")
	panic(err)
}
```

大多数程序只有在出现意料之外的错误时才会出现panic。你应该考虑程序可能遇到的所有错误（例如文件丢失或格式错误的数据），并使用error值来处理这些错误。