package main

import (
	"context"
	"fmt"
	"rpc_start/rpc11_grpc_interpretor/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {

	interceptor := func(ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		//获取时间差
		fmt.Println("耗时:%s", time.Since(start))
		return err
	}
	//方式一
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial(":7777", grpc.WithInsecure(), opt)

	//方式二
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial(":7777", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "wyang"})
	if err != nil {
		return
	}

	fmt.Println(r.Message)

}
