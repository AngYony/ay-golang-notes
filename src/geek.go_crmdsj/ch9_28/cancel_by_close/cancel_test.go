package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case wy := <-cancelChan:
		// 如果往取消队列中插入了值，就说明取消操作
		fmt.Println("sssss", wy)
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	// 向队列中插入一个空结构
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				} else {
					fmt.Println(i)
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")

		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 10)
}
