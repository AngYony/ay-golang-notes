package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(a int) {
			for {
				arr[a]++
				runtime.Gosched() //手动交出控制权，让别的goroutine也有机会运行，不推荐使用
				fmt.Printf("Hello from goroutine %d \n", a)
			}
		}(i)
	}

	time.Sleep(time.Millisecond)

}
