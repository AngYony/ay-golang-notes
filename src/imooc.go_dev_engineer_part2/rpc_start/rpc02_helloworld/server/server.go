package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}

func main() {

	// 1. 实例化一个Server
	// 设置监听的是一个tcp的端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 注册处理逻辑 handler
	// 将struct注册到RPC中，指定Service的名称，可以任意命名，最终会作为包名被客户端调用
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		log.Fatal(err)
	}

	// 3. 启动服务
	conn, err := listen.Accept() // 当一个新的连接进来之后，就有一个套接字（socket）
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
