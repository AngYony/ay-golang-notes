package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d 接收值：%d \n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	//定义一个定时器
	tm := time.After(10 * time.Second)
	//每一秒钟查看队列中积压的元素个数
	tick := time.Tick(time.Second) //指定间隔时间重复执行的定时器
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		//每次select的时间超过800毫秒
		case <-time.After(800 * time.Millisecond):
			//相邻的两个请求之间timeout
			fmt.Println("超时")
		case <-tick:
			fmt.Println("队列长度：", len(values))
		case <-tm:
			//总的时间
			//定时器channel获取到了值
			fmt.Println("再见")
			return

		}
	}

}
