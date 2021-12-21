package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环的接收客户端发送的数据
	defer conn.Close() // 关闭conn
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)

		// fmt.Printf("服务器在等待客户端【%s】发送信息...", conn.RemoteAddr().String())
		// 等待客户端通过conn发送信息
		// 如果客户端没有write【发送】，那么协程就阻塞在这里
		n, err := conn.Read(buf) // 从conn读取
		// if err == io.EOF {
		if err != nil {
			fmt.Println("客户端已退出")
			return
		}
		// 显示客户端发送的内容到服务器的终端
		// 由于客户端发送的时候已经按行发送，所以这里只用Print就行，不需要Println
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听端口....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	defer listen.Close()

	// 循环等待客户端的连接
	for {
		// 等待客户端连接
		fmt.Println("等待客户端来连接。。。")
		conn, err := listen.Accept() // 有一个Accept代表一个客户端，每次有新的客户端都会有新的Conn

		if err != nil {
			fmt.Println("Accept() 失败")
			// 此处不需要return，一旦return就拒绝了所有的连接了
		} else {
			fmt.Println("Accept()成功,客户端IP：", conn.RemoteAddr().String())
		}

		// 这里准备开启一个协程，为客户端服务
		go process(conn)

	}

}
