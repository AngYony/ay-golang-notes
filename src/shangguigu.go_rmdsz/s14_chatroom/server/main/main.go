package main

import (
	"fmt"
	"net"
	"shangguigu.Go_rmdsz/s14_chatroom/server/model"
	"time"
)

func process(conn net.Conn) {
	// 读客户端发送的信息
	defer conn.Close()

	// 调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程出错,err=", err)
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("59.110.216.174:6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("服务器在8889端口监听...")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return

	}

	defer listen.Close()
	// 一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}
