package anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("结果来自于 id:%d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch

}

func TestFirstResponse(t *testing.T) {
	// 输出系统中的协程数
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}
