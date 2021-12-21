package main

import (
	"context"
	"fmt"
	"net"
	"rpc_start/rpc10_metadata_test/proto"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("获取 metadata 错误")
	}

	//如果直接遍历metadata，会将header部分数据也获取到
	for key, val := range md {
		fmt.Println(key, val)
	}

	//如果想要获取指定key的值
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		//注意：metadata的每个key对应的value是一个切片类型，因此想要获取具体值，可以使用range进行读取
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

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
