//grpc中的错误处理
package main

import (
	"context"
	"net"
	"rpc_start/rpc13_error_test/proto"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	//返回状态码的错误信息
	//return nil, status.Errorf(codes.NotFound, "未找到:%s", request.Name)

	time.Sleep(time.Second * 5) //模拟服务端处理超时
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
