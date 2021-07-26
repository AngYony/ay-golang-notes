package main

import (
	"fmt"
	"net"
	"rpc_start/rpc08_stream_grpc_test/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":2222"

type server struct {
}

//服务器端流模式
func (s *server) GetStream(req *proto.StreamRequestData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResponseData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 20 {
			break
		}
	}

	return nil
}

//客户端流模式
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

//双向流模式
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResponseData{Data: "服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

//rpc (StreamRequestData) returns (stream StreamResponseData);
//
////客户端流模式
//rpc PutStream(stream StreamRequestData) returns ( StreamResponseData);
//
////双向流模式
//rpc AllStream(stream StreamRequestData) returns (stream StreamResponseData);

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	//创建grpcserver
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	err = s.Serve(listen)

}
