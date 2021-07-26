package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_start/rpc05_new_helloworld/handler"
	"rpc_start/rpc05_new_helloworld/server_proxy"
)

func main() {

	//1. 实例化一个Server
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	//2. 注册处理逻辑 handler
	err = server_proxy.RegisterHelloService(&handler.NewHelloService{})

	if err != nil {
		log.Fatal(err)
	}
	for {
		//3. 启动服务
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
