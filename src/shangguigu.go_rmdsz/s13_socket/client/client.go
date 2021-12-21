package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败:", err)
		return
	}

	fmt.Println("客户端连接成功。。。", conn)

	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入（终端）

	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败")
		}
		line = strings.Trim(line, "\r\n")
		line = strings.TrimSpace(line)
		if line == "exit" {
			fmt.Println("退出")
			break
		}

		// 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write()出现异常")
		}
	}
	// fmt.Printf("客户端发送了%d 字节的数据，并退出\n", n)
}
