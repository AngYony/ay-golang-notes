//grpc中的错误处理和超时
package main

import (
	"context"
	"fmt"
	"rpc_start/rpc13_error_test/proto"
	"time"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	//设置超时，一点服务端超时，将会接收到超时相关的error信息
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "wyang"})
	if err != nil {

		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())

	}

	//fmt.Println(r.Message)
}
