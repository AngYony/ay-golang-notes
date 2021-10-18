package main

import (
	"context"
	"fmt"
	"log"
	"rpc_start/grpclb_test/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	// 拨号
	conn, err := grpc.Dial(
		"consul://192.168.171.223:8500/user-srv?wait=14s&tag=wy",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userSrvClient := proto.NewUserClient(conn)
	for i := 0; i < 10; i++ {

		list, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
			Pn:    1,
			PSize: 2,
		})

		if err != nil {
			panic(nil)
		}
		for index, data := range list.Data {
			fmt.Println(index, data)
		}
	}
	// for i := 0; i < 10; i++ {
	// 	userSrvClient := proto.NewUserClient(conn)
	// 	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
	// 		Pn:    1,
	// 		PSize: 2,
	// 	})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for index, data := range rsp.Data {
	// 		fmt.Println(index, data)
	// 	}
	// }

}
