package main

import (
	"fmt"
	hello "rpc_start/rpc06_protobuf/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	req := hello.HelloRequest{
		Name: "wy",
	}

	rsp, _ := proto.Marshal(&req)
	fmt.Println(rsp)
}
