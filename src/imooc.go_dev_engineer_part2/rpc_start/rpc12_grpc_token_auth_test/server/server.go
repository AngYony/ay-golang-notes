//Server端拦截器的使用示例
package main

import (
	"context"
	"fmt"
	"net"
	"rpc_start/rpc12_grpc_token_auth_test/proto"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

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

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无Token认证信息")
		}

		var appid string
		var appkey string

		if val1, ok := md["appid"]; ok {
			appid = val1[0]
		}

		if val1, ok := md["appkey"]; ok {
			appkey = val1[0]
		}

		if appid != "wy" || appkey != "MMM" {
			return resp, status.Error(codes.Unauthenticated, "id或Key无效")
		}

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
