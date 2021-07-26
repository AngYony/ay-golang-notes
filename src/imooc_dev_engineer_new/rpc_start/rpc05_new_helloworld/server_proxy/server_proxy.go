package server_proxy

import (
	"net/rpc"
	hanlder "rpc_start/rpc05_new_helloworld/handler"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
