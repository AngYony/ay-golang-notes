package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/proto"
	"net"

	"google.golang.org/grpc"
)

// 使用说明：
/*
cd 当前目录
go build main.go
main.exe -h
main.exe -port 50053

*/

func main() {
	// 接收用户输入的命令参数，来监听对应的ip和端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
