package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

//将channel作为参数
func doWork(id int, w worker) {
	//从channel中读数
	for n := range w.in {
		fmt.Printf("Worker %d 接收值：%d \n", id, n)
		//调用外部定义的done函数
		w.done()
	}
}

//返回channel，这个channel在外部只能发数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	//创建channel
	w := worker{
		in: make(chan int),
		//定义done函数将要执行的操作
		done: func() {
			wg.Done()
		},
	}
	//定义channel取数goroutine
	go doWork(id, w)
	return w
}

//具备方向指向的channel的使用
func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20) //添加任务数量

	for i, worker := range workers {
		//向每个channel发送数据
		worker.in <- 'a' + i
	}
	//进行第二批发数
	for i, worker := range workers {
		//向每个channel发送数据
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	wg.Wait()

}

func main() {
	chanDemo()
}
