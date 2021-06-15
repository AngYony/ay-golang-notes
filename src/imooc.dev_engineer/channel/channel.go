package main

import (
	"fmt"
	"time"
)

//将channel作为参数
func worker(id int, c <-chan int) {
	for n := range c {
		fmt.Printf("Worker %d 接收值：%d \n", id, n)
	}

	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d 接收值：%d \n", id, n)
	}
}

//返回channel，这个channel在外部只能发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	//对channel持续取值
	go func() {
		for {
			fmt.Printf("Worker %d 接收值：%c \n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		//启动goroutine
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		//向每个channel发送数据
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		//向每个channel发送数据
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

//具备方向指向的channel的使用
func chanDemo2() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		//向每个channel发送数据
		channels[i] <- 'a' + i

		//尝试收数据将会报错
		//n:=<- channels[i]
	}

	for i := 0; i < 10; i++ {
		//向每个channel发送数据
		channels[i] <- 'A' + i
	}
}

//定义可缓冲的channel
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)

}

func main() {

	//bufferedChannel()
	channelClose()
	time.Sleep(time.Second)
}
