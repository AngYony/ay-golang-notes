//Server端拦截器的使用示例
package main

import (
	"context"
	"fmt"
	"net"
	"rpc_start/rpc11_grpc_interpretor/proto"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context,
	request *proto.HelloRequest) (*proto.HelloReply, error) {

	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {

	interceptor := func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		fmt.Println("接收到了一个新的请求")
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成") //这种可以用来统计请求使用了多长时间

		return res, err

	}

	opt := grpc.UnaryInterceptor(interceptor)

	g := grpc.NewServer(opt)

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
