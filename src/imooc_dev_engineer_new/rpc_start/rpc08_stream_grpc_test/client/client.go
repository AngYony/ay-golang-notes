package main

import (
	"context"
	"fmt"
	"rpc_start/rpc08_stream_grpc_test/proto"

	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":2222"

func main() {

	//初始化连接

	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//服务器端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamRequestData{Data: "学习"})
	for {
		recv, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Data)

	}

	//客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamRequestData{
			Data: fmt.Sprintf("哈哈%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式

	allStr, _ := c.AllStream(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}

	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamRequestData{Data: "FFFFFFFF"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
