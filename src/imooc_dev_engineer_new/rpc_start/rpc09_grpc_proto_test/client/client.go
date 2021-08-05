package main

import (
	"context"
	"fmt"
	"rpc_start/rpc09_grpc_proto_test/proto"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(),
		&proto.HelloRequest{
			Name: "wyang",
			G:    proto.Gender_MALE,
			Mp: map[string]string{
				"name": "张三",
				"work": "打杂的",
			},
			AddTime: timestamppb.New(time.Now()),
		})
	if err != nil {
		return
	}

	fmt.Println(r.Message)

}
