# Go-goroutine和channel

预想问题

goroutine和chanel之间的关系？

goroutine是什么？

chanel是什么？

channel允许goroutine互相发送数据并同步，这样一个goroutine就不会领先于另一个goroutine。？



## goroutine

并发：一次不止完成一件事。

并行：同时运行多个任务，它是并发的一种形式。

goroutine可以让程序同时处理几个不同的任务。

在Go中，并发任务 称为goroutine，类似于C#中的多线程。但是goroutine比线程需要更少的计算机内存，启动和停止的时间更少，这意味着你可以同时运行更多的goroutine。

使用go语句来启动一个新的goroutine。

```
go 函数名()
```

每个Go程序的main函数都是使用goroutine启动的，称为main goroutine，因此每个Go程序至少运行一个goroutine。（对于goroutine，可以将其想象成主线程，每一个程序至少由一个线程运行）

### goroutine的使用

```go
func a() {
	//使用循环打印500个字母a
	for i := 0; i < 500; i++ {
		fmt.Print("a")
	}
}

func b() {
	//使用循环打印500个字母b
	for i := 0; i < 500; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	//暂停main goroutine 3秒
	time.Sleep(time.Second * 3)

	fmt.Println("结束")
}
```

运行上述的代码，将会混合输出字母“a”和字母“b”，多运行几次，字母“b”比字母“a”还先输出。

在C#中，一旦主线程结束了，其他线程也会停止运行。和C#一样，==Go程序在main goroutine（调用main函数的goroutine）结束后立即停止运行，即使其他goroutine仍在运行。==因此上述代码使用了休眠3秒钟，以便等待其他goroutine执行完毕。

使用go语句启动一个新的goroutine时，其实是一个异步运行的过程。因此多个go语句启动多个goroutine，就能够实现并行处理。

需要特别注意的是：==Go不能保证goroutine按照调用go语句的先后顺序依次运行==，上述代码虽然`go b()`在`go a()`之后被调用，但是输出的结果，却可能先输出的是字母“b”。==因此，goroutine在内部是按照最有效的方式运行的，Go不能保证何时在goroutine之间切换，或者切换多长时间==。例如上述代码中，需要从main goroutine切换到`go a()`的goroutine或者`go b()`的goroutine，到底先切换到哪一个，完全取决于goroutine本身的运行方式，如果goroutine运行的顺序很重要，那么必须使用channel来同步它们。

由于go语句是按照异步的方式运行的，因此不能直接使用函数返回值。

例如，下述操作是不被允许的，将会报编译错误：

```go
wy:=go runSize()
```

原因很容易理解，就像C#中的异步返回方法需要使用await关键字一样，而go语句本身就是异步执行的，直接按照上述方式接收返回值，是不被允许的，解决办法是使用channel。



## channel

channel可以将值从一个goroutine发送到另一个goroutine，并且可以确保在接收的goroutine尝试使用该值之前，发送的goroutine已经发送了该值。

chanel用于一个goroutine到另一个goroutine的通信。

### 创建channel

每个channel只携带特定类型的值。使用chan关键字来声明包含channel的变量，并指定channel将携带的值的类型。

语法示例：

```go
var myChannel chan float64
```

上述只是声明，要实际创建channel，需要调用内置的make函数，并为make()函数传入要创建的channel的类型（应该与要赋值给它的变量的类型相同）。

```go
var myChannel chan float64
myChannel = make(chan float64)
```

大多数情况下，直接使用一个短变量声明并创建channel：

```go
myChannel := make(chan float64)
```

### 使用chanel发送和接收值

通过使用`<-`运算符来发送和接收值，不同的是`<-`在channel的前后位置不同。

发送值，使用`<-`运算符，位于channel的右侧，从发送的值指向发送该值的channel：

```go
myChannel <- 3.14
```

接收值，`<-`位于channel的左侧：

```go
<- myChannel
```

记忆技巧：程序从右往左执行，因此位于channel右侧的就是发送值，左侧就是取值。

创建一个以channel作为参数的函数，使用channel发送值：

```go
func greeting(myChannel chan string) {
	myChannel <- "Hi" //通过channel发送一个值
}
```

使用channel接收值：

```go
func main() {
	//创建一个新的channel
	myChannel := make(chan string)
	//将channel传递给新goroutine中运行的函数
	go greeting(myChannel)
	//从channel中接收值
	chv := <-myChannel
	fmt.Println(chv)
}
```



### 有缓冲的channel

在创建channel时，可以通过给make传递第二个参数来创建有缓冲的channel，该参数包含channel应该能够在其缓冲区中保存的值的数量。

```go
channel := make(chan string,3)
```

有缓冲的channel可以在导致发送的goroutine阻塞之前保存一定数量的值。在适当的情况下，这可以提高程序的性能。

当goroutine通过channel发送一个值时，该值被添加到缓冲区中。发送的goroutine将继续运行，而不被阻塞。

==发送的goroutine可以继续在channel上发送值，直到缓冲区被填满；只有这时，额外的发送操作才会导致goroutine阻塞。==

当另一个goroutine从channel接收一个值时，它从缓冲区中提取最早添加的值。

额外的接收操作将继续清空缓冲区，而额外的发送操作将填充缓冲区。



## goroutine 与 channel 之间的同步

channel可以确保发送的goroutine 在接收channel尝试使用该值之前已经发送了该值。

==无论是发送值还是接收值，channel都会阻塞当前的goroutine==：

- ==发送操作阻塞发送goroutine，直到另一个goroutine在同一channel上执行了接收操作。==
- ==接收操作阻塞接收goroutine，直到另一个goroutine在同一channel上执行了发送操作。==

channel通过blocking（阻塞）——暂停当前goroutine中的所有进一步操作来实现这一点。

例如：

```go
func abc(mych chan string) {
	mych <- "a"
	mych <- "b"
	mych <- "c"
}

func def(mych chan string) {
	mych <- "d"
	mych <- "e"
	mych <- "f"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go abc(ch1) //发送abc
	go def(ch2) //发送def

	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
	fmt.Print(<-ch1)
	fmt.Print(<-ch2)
}
```

在上述程序中，abc()和def()均以异步的方式并发执行。在abc的goroutine中，每次向channel发送一个值，都会阻塞abc的goroutine，直到有其他的goroutine接收到这个channel的值为止。def的goroutine也是同样如此。同样，当main goroutine接收每个channel参数时，也会阻塞main的goroutine，直到abc或def的goroutine执行发送操作，之前已经说过，channel可以确保发送的goroutine 在接收channel尝试使用该值之前已经发送了该值，main goroutine成为abc goroutine和def goroutine的协调器，只有当它准备读取它们发送的值时，才允许它们继续。因此上述程序将会按照获取值的顺序输出：adbecf。

综合示例：

```go
func reportNap(name string, delay int) {
	//每一秒打印一个通知，说还在休眠
	for i := 0; i < delay; i++ {
		fmt.Println(name, "正在休眠")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "休眠结束")
}

func send(myChannel chan string) {
	//休眠2秒
	reportNap("发送前休眠2秒", 2)
	fmt.Println("***sending value a***")
	myChannel <- "a" //发送值，阻塞当前goroutine，直到其他goroutine接收该channel的值
	fmt.Println("***sending value b***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel) //将以异步方式运行
	//休眠5秒
	reportNap("5秒休眠", 5)
	//直到5秒之后，才接收值，从而解除send goroutine中的阻塞
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
```

输出：

```go
5秒休眠 正在休眠
发送前休眠2秒 正在休眠
5秒休眠 正在休眠
发送前休眠2秒 正在休眠
发送前休眠2秒 休眠结束
***sending value a***
5秒休眠 正在休眠
5秒休眠 正在休眠
5秒休眠 正在休眠
5秒休眠 休眠结束
a
***sending value b***
b
```

在main goroutine中获取channel的值时，一定要保证在获取语句之前，已经使用了go语句启动新的goroutine，并为channel发送了值，否则接收channel所在的main goroutine将会一直被阻塞，会引发异常。

例如，下述程序将会引发异常：

```go
func greeting2(myChannel chan string) {
	myChannel <- "hi" //发送操作会导致该goroutine阻塞
}
func main() {
	myChannel := make(chan string)
	//在main goroutine中发送值，会阻塞main goroutine，直到其他goroutine获取值
	myChannel <- "你看"
    //由于阻塞，不会被输出
	fmt.Println("被阻塞了")
	//go语句由于上述的阻塞，不会被执行，不会出现接收的goroutine，因此引发异常
	go greeting2(myChannel)
	fmt.Println(<-myChannel)
}
```

输出错误：

```
fatal error: all goroutines are asleep - deadlock!
...
```

