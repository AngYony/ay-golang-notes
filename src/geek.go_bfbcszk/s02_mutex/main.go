package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Name string

	mu    sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计算器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	var counter Counter // 不需要显示初始化Mutex即可使用
	// 使用 WaitGroup 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Count())
}
