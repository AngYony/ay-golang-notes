package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	//1. 建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	err = client.Call("HelloService.Hello", "wyang", &reply)

	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

}

//{"method":"HelloService.Hello","params":["wyang"],"id":0}
