package main

import (
	"context"
	"fmt"
	"rpc_start/rpc07_grpc_test/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "wyang"})
	if err != nil {
		return
	}

	fmt.Println(r.Message)
}
