package main

import "sync"

//声明互斥锁
var mu sync.Mutex

func main() {
	mu.Lock()         //对互斥锁执行上锁操作
	defer mu.Unlock() //在函数返回之前解锁互斥锁
	//在函数返回之前，互斥锁始终处于锁定状态
}
