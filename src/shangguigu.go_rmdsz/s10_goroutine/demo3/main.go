package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("写入数据：", i)

	}
	// 关闭管道
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		// 判断是否关闭，如果关闭，OK=false
		if !ok {
			break

		}
		fmt.Printf("读取到数据：%d \n", v)
	}

	// 读取完成之后，发送任务完成状态
	exitChan <- true
	close(exitChan)
}

func main() {
	intchan := make(chan int, 10)
	existchan := make(chan bool, 1)

	go writeData(intchan)
	go readData(intchan, existchan)
	<-existchan

}
