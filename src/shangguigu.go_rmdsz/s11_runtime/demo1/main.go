package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 获取CPU逻辑核心数
	cpuNum := runtime.NumCPU()
	fmt.Println("CPU逻辑核心数：", cpuNum)

	// 设置使用多少个CPU核心数，不需要显式设置
	runtime.GOMAXPROCS(cpuNum - 2)
}
