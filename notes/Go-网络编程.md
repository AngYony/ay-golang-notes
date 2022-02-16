# Go - 网络编程



服务端处理流程

- 监听端口，例如：8888
- 接收客户端的TCP连接，建立客户端和服务器端的连接
- 创建goroutine，处理该连接的请求



客户端处理流程

- 建立与服务端的连接
- 发送请求数据，接收服务器端返回的结果数据
- 关闭连接



## net 包

提供了网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域Socket。





技巧：

验证并连接到某个端口，可以使用telnet命令。

例如，cmd中输入：`telnet www.baiu.com 80`，将会进入到telnet工具中。

退出按下快捷键`ctrl+]`，然后输入`quit`命令回到cmd控制台。



## 综合示例

一对一发送消息程序。

服务端：

```go
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
```

客户端:

```go
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
```

