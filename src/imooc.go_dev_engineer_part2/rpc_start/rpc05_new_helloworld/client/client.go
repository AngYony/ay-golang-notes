package main

import (
	"fmt"
	"rpc_start/rpc05_new_helloworld/client_proxy"
)

func main() {

	//1. 建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	//var reply *string //如果直接这么使用会报错，原因是reply为nil，没有具体地址指针信息
	var reply *string = new(string) //new的作用：在内存分配空间，并把指针地址赋值给变量
	err := client.Hello("wyang", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)

}
