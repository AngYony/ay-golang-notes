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

	//1. 实例化一个Server
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	//2. 注册处理逻辑 handler

	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		log.Fatal(err)
	}

	//3. 启动服务
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
