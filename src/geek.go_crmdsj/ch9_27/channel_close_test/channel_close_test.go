package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 生产完关闭 channel
		close(ch)
		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			data, ok := <-ch
			if ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()

	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	// 调用生产者
	dataProducer(ch, &wg)

	wg.Add(1)
	// 调用消费者
	dataReceiver(ch, &wg)

	wg.Wait()
}
