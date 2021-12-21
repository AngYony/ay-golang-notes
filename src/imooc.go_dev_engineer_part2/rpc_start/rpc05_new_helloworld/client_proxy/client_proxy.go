package client_proxy

import (
	"log"
	"net/rpc"
	hanlder "rpc_start/rpc05_new_helloworld/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) HelloServiceStub {
	client, err := rpc.Dial(protol, address)
	if err != nil {
		log.Fatal(err)
	}
	return HelloServiceStub{client}
}
func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Client.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
