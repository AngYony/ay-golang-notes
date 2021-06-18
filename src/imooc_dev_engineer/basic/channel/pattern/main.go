//并发模型一
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("%s, message %d", name, i)
			i++
		}
	}()
	return c
}

//任务中断
func msgGen2(name string, done chan struct{}) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("%s, message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second) //模拟2秒清理完
				//双向通信
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}

func fanIn2(chs ...<-chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in <-chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

//并行模式二：select
func fanInBySelect(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()
	return c
}

//非阻塞式等待
func nonBlockingWait(c <-chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c <-chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})

	m1 := msgGen2("s1", done)

	for i := 0; i < 5; i++ {

		if m, ok := timeoutWait(m1, time.Second); ok {
			//if m, ok := nonBlockingWait(m1); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}

	}
	//双向通信
	done <- struct{}{}
	<-done
}

func main2() {
	m1 := msgGen("s1")
	m2 := msgGen("s2")

	m := fanIn2(m1, m2)

	for {
		fmt.Println(<-m)

	}
}
