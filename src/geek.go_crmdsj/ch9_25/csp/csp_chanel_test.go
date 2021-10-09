package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "返回 service 的执行结果"
}

func otherTask() {
	fmt.Println("开始执行另一个任务")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("另一个任务执行完成")

}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)

	// retCh := make(chan string, 1)
	go func() {
		ret := service()
		// todo:最终输出的结果，为什么是先输出"service 获取到了结果"，后输出"返回 service 的执行结果"
		fmt.Println("service 获取到了结果")
		retCh <- ret
		fmt.Println("service 结果放入到了 chanel 中并退出")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)

}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)

	case <-time.After(time.Millisecond * 100):
		// 超时实现
		fmt.Println("time out")

	}
}

func wy() {
	fmt.Println("测试呀")
	time.Sleep(time.Millisecond * 10)
}
func TestWy(t *testing.T) {
	go func() {
		wy()
		fmt.Println("aaa")

	}()
	time.Sleep(time.Second * 1)
}
