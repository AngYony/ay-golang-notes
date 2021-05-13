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

每次使用关键字<code>go</code>都会产生一个新的goroutine。从表面上来看，所有goroutine似乎都在同时运行，但由于计算机通常只具有有限数量的处理单元，因此从技术上说，这些goroutine并不是真的在同时运行。

实际上，计算机的处理器通常会使用一种名为分时的技术，在多个goroutine上面轮流花费一些时间。因为分时的具体实施细节通常只有Go运行时、操作系统和处理器会知道，所以我们在使用goroutine的时候，应该假设不同goroutine中的各项操作将以任意顺序执行。

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



## channel（通道）

channel可以将值从一个goroutine发送到另一个goroutine，并且可以确保在接收的goroutine尝试使用该值之前，发送的goroutine已经发送了该值。

chanel用于一个goroutine到另一个goroutine的通信。

跟Go中的其他类型一样，可以将通道用作变量、传递至函数、存储在结构中等几乎任何事情。

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

除goroutine本身占用的少量内存之外，被阻塞的goroutine并不消耗任何资源。goroutine会静静地停在那里，等待导致它阻塞的事情发生，然后解除阻塞。

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



## 使用select处理多个 channel

下面是一个未使用select的示例：

```go
func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " 睡眠中 ...")
	c <- id //发送值
}

func main() {

	c := make(chan int)
    //同时启动5个goroutine
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	fmt.Println("已完成goroutine的全部启动")
	
	for i := 0; i < 5; i++ {
		gopherId := <-c //从通道中取值
		fmt.Println("gopher ", gopherId, "完成睡眠")
	}
}
```

执行上述代码输出如下结果：

```
已完成goroutine的全部启动
... 0  睡眠中 ...
... 4  睡眠中 ...
... 3  睡眠中 ...
gopher  0 完成睡眠
gopher  4 完成睡眠
gopher  3 完成睡眠
... 2  睡眠中 ...
gopher  2 完成睡眠
... 1  睡眠中 ...
gopher  1 完成睡眠
```

<code>select</code>语句跟<code>switch</code>语句有点儿相似，该语句包含的每个<code>case</code>分支都持有一个针对通道的接收或发送操作。<code>select</code>会等待直到某个分支的操作就绪，然后执行该操作及其关联的分支语句，它就像是在同时监控两个通道，并在发现其中一个通道出现情况时采取行动。

使用select示例：

```go
func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " 睡眠中 ...")
	c <- id //发送值
}

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	fmt.Println("已完成goroutine的全部启动")

	
	//time.After函数返回一个通道
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c: //从通道中取值
			fmt.Println("gopher ", gopherId, "完成睡眠")
		case <-timeout:
			fmt.Println("等待超时")
			return
		}
	}
}
```

由于每个goroutine中都休眠了3秒，所以会直接输出“等待超时”：

```
已完成goroutine的全部启动
等待超时
```

上述代码使用了time.After函数来创建超时通道。

time.After函数会返回一个channel，该channel会在经过特定时间之后，接收到一个值（发送该值的goroutine是Go运行时的其中一部分）。

<code>select</code>语句在不包含任何分支的情况下将永远地等待下去。当你启动多个goroutine并且打算让它们无限期地运行下去的时候，就可以用这个方法来阻止<code>main</code>函数返回。

注意：即使程序已经停止等待goroutine，但只要<code>main</code>函数还没返回，仍在运行的goroutine就会继续占用内存。所以在情况允许的情况下，我们还是应该尽量结束无用的goroutine。

只包含一个分支的select语句实际上跟直接执行通道操作的效果是一样的。

### nil通道

对值为nil的通道执行发送或接收操作并不会引发惊恐，但是会导致操作永久阻塞，就好像遇到了一个从来没有接收或者发送过任何值的通道一样。但如果你尝试对值为nil的通道执行close函数，将会引发惊恐。

初看上去，值为nil的通道似乎没什么用处，但事实恰恰相反。例如，对于一个包含<code>select</code>语句的循环，如果我们不希望程序在每次循环的时候都等待<code>select</code>语句涉及的所有通道，那么可以先将某些通道设置为nil，等到待发送的值准备就绪之后，再为通道变量赋予一个非 nil 值并执行实际的发送操作。



## 关闭通道

Go允许在没有值可供发送的情况下通过close函数关闭通道。

通道被关闭之后将无法写入任何值，如果尝试写入值将会引发惊恐。尝试读取已被关闭的通道将会获得一个与通道类型对应的零值。

注意：　当心！如果你在循环里面读取一个已关闭的通道，并且没有检查该通道是否已经关闭，那么这个循环将一直运转下去，并耗费大量的处理器时间。为了避免这种情况发生，请务必对那些可能会被关闭的通道做相应的检查。

执行下述代码获悉通道是否已经关闭：

```go
v, ok := <-c
```

如果第二个变量ok的值为false，说明通道已经被关闭。

因为“从通道里面读取值，直到它被关闭为止”这种模式实在是太常用了，所以Go为此提供了一种快捷方式。通过在<code>range</code>语句里面使用通道，程序可以在通道被关闭之前，一直从通道里面读取值。即读取数据时，不用每次都判断通道是否已关闭。

```go
//形参是两个通道
func filterGopher(upstream, downstream chan string) {
    //使用range语句，迭代通道中的值
    for item := range upstream {
        if !strings.Contains(item, "bad") {
            downstream <- item
        }
    }
    close(downstream)
}
```



## 并发状态

### 互斥锁（Mutex）

goroutine可以通过互斥锁阻止其他goroutine在同一时间进行某些操作。

互斥锁具有Lock和Unlock两个方法。

如果有goroutine尝试在互斥锁已经锁定的情况下调用Lock方法，那么它就需要等到解除锁定之后才能够再次上锁。

为了正确地使用互斥锁，需要确保所有访问共享值的代码必须先锁定互斥锁，然后才能执行所需的操作，并且在操作完成之后必须解除互斥锁。任何不遵循这一模式的代码都可能会引发竞态条件。基于上述原因，互斥锁在绝大多数情况下只会在包的内部使用。包会通过互斥锁保护指定的内容，并将相应的<code>Lock</code>和<code>Unlock</code>调用巧妙地隐藏在方法和函数的背后。

和channel不一样，互斥锁并未内置在Go语言当中，而是通过sync包提供。

```go
//声明互斥锁，不需要对其实施初始化，其零值就是一个未上锁的互斥锁
var mu sync.Mutex

func main() {
	mu.Lock()         //对互斥锁执行上锁操作
	defer mu.Unlock() //在函数返回之前解锁互斥锁
	//在函数返回之前，互斥锁始终处于锁定状态
}
```

将sync.Mutex用作结构成员的做法是一种常见的模式：

```go
// Visited用于记录网页是否被访问过
// 它的方法可以在多个goroutine中并发使用
type Visited struct {
    // mu 用于保护 visited 映射
    mu      sync.Mutex  	//声明一个互斥锁
    visited map[string]int  
}
```

对应的互斥锁的使用：

```go
// VisitLink 会记录本次针对给定网址的访问，然后返回更新之后的链接统计值
func (v *Visited) VisitLink(url string) int {
    v.mu.Lock()  			//锁定互斥锁
    defer v.mu.Unlock()  	//确保锁定会在之后解除
    count := v.visited[url]
    count++
    v.visited[url] = count 	//更新映射
    return count
}
```

程序在锁定之后需要执行的操作越多，我们越要小心。如果一个goroutine在锁定互斥锁之后因为某些事情而被阻塞，那么想要取得互斥锁的其他goroutine就可能会被耽搁很长一段时间。更严重的是，如果持有互斥锁的goroutine因为某些原因而尝试锁定同一个互斥锁，那么就会引发死锁——正在尝试执行加锁操作的goroutine将永远无法解除已经被锁定的互斥锁，最终导致<code>Lock</code>调用被永久阻塞。

为了保证互斥锁的使用安全，我们必须遵守以下规则：

- 尽可能地简化互斥锁保护的代码。
- 对每一份共享状态只使用一个互斥锁。



## 工作进程（worker）

一直存在并且独立运行的goroutine称为工作进程（worker）。

工作进程通常会被写成包含<code>select</code>语句的<code>for</code>循环。只要工作进程在运行，循环就会继续下去，而<code>select</code>则会等待某些事情发生。使用长时间运行的goroutine可以实现带有select循环的工作进程。

以下是一个没有任何实际用途的工作进程的函数框架：

```go
func worker() {
    for {
        select {
        // 在此处等待通道
        }
    } 
}
```

启动工作进程：

```go
go worker()
```



在Go语言中，通道常常被看作是实现细节，所以一般都会把通道隐藏在方法的后面。

综合示例：

```go
package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

type command int

//模拟两个command常量
const (
	right = command(0)
	left  = command(1)
)

type RoverDriver struct {
	//定义一个发送命令的通道
	commandc chan command
}

//定义向左的方法
func (r *RoverDriver) Left() {
	//向通道发送left命令值
	r.commandc <- left
}

//定义向右的方法
func (r *RoverDriver) Right() {
	//向通道发送right命令值
	r.commandc <- right
}

//定义结构的drive方法，能够访问RoverDriver的任何成员
func (r *RoverDriver) drive() {
	//当前位置初始值
	pos := image.Point{X: 0, Y: 0}
	//当前方向
	direction := image.Point{X: 1, Y: 0}

	updateInterval := 250 * time.Millisecond
	//创建初始计时器通道
	nextMove := time.After(updateInterval)

	for {
		select {
		//等待接收来自命令通道的命令
		case c := <-r.commandc:
			//判断命令的值，执行不同的分支操作
			switch c {
			//向右转
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				//向左转
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("new direction %v", direction)

		case <-nextMove: //从通道中取到值后将会击发计时器
			pos = pos.Add(direction)
			fmt.Println("当前位置：", pos)

			//为下一次事件循环创建新的计时器通道
			nextMove = time.After(updateInterval)
		}
	}
}

//创建通道并启动工作进程
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func main() {
	r := NewRoverDriver()
	//此处休眠3秒，将始终触发time.after通道，将会连续输出“当前位置”信息，直到3秒结束
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
```



