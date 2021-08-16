//使用全局拦截器实现认证功能
//方式二：定义一个新的类型，实现PerRPCCredentials接口
package main

import (
	"context"
	"fmt"
	"rpc_start/rpc12_grpc_token_auth_test/proto"

	"google.golang.org/grpc"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "wy",
		"appkey": "MMM",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	//interceptor := func(ctx context.Context,
	//	method string,
	//	req, reply interface{},
	//	cc *grpc.ClientConn,
	//	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//
	//	md := metadata.New(map[string]string{
	//		"appid":  "wy",
	//		"appkey": "MMM",
	//	})
	//
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	//获取时间差
	//	fmt.Println("耗时:%s", time.Since(start))
	//	return err
	//}
	//方式一
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial(":7777", grpc.WithInsecure(), opt)

	//grpc.WithPerRPCCredentials(customCredential{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
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
