package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// 1. 建立连接
	// 在本地的1234端口，发起一个tcp的连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	// var reply *string //如果直接这么使用会报错，原因是reply为nil，没有具体地址指针信息
	var reply *string = new(string) // new的作用：在内存分配空间，并把指针地址赋值给变量
	// 调用注册到RPC中的服务
	err = client.Call("HelloService.Hello", "wyang", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)

}
