package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享资源
var myMap = make(map[int]int, 10)

// 定义全局互斥锁
var lock sync.Mutex

// 计算1+2+...n的值
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res += i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go factorial(i)
	}

	time.Sleep(time.Second * 5)

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
