package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80; i++ {
		intChan <- i
	}
	// 写入完之后要随手关闭管道
	close(intChan)
}

func primeNum(chanNum int, intChan chan int, primeChan chan string, exitChan chan bool) {

	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag := true
		// 判断是否素数：除了1和自己都不能被整除
		for i := 2; i < num; i++ {
			// 不是素数
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- fmt.Sprintf("管道%d计算出的素数：%d", chanNum, num)
		}
	}
	fmt.Println("有一个协程取不到数据已退出")
	exitChan <- true
	// 此处不能关闭，因为有可能别的协程在写入数据

}

func main() {
	intChan := make(chan int, 10)
	primeChan := make(chan string, 10) // 存放结果
	exitChan := make(chan bool, 4)     // 4个

	// 开启协程，放入数据
	go putNum(intChan)

	// 开启多个协程，取数并计算是否是素数，并将结果放入到结果管道中

	// 启动4个协程计算素数
	for i := 1; i <= 4; i++ {
		go primeNum(i, intChan, primeChan, exitChan)
	}

	// 检测是否上述4个协程都执行完成，执行完成之后，需要关闭管道
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()

	// 启动一个新的协程用于输出结果
	// 输出素数
	for v := range primeChan {
		fmt.Println(v)
	}

}
