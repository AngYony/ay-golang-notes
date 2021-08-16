package main

import (
	"context"
	"fmt"
	"rpc_start/rpc10_metadata_test/proto"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format(""))

	md2 := metadata.New(map[string]string{
		"name": "张三",
		"pwd":  "abcd",
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md2)

	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "wyang"})
	if err != nil {
		return
	}

	fmt.Println(r.Message)
}
