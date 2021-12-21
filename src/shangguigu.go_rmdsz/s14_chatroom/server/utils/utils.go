package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.Go_rmdsz/s14_chatroom/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // 数据传输时，使用的缓冲
}

// 读取客户端消息
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 1024*4) // 能够接收到的最大字节数
	fmt.Println("读取客户端发送的数据...")
	// 返回的n表示实际上读了多少字节
	// 如果客户端关闭了Conn，此处的Read就不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}

	// 将字节数组转换为uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	// 从conn套接字中读取pkglen个字节扔到buf中去
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail,err=", err)
		return
	}
	return
}

// 发送消息给客户端
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	// var buf [4]byte // 因为uint32数字存储需要4个字节，所以这里使用的是4个长度的字节数组
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[0:4])

	if n != 4 || err != nil {

		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	// 发送消息本身
	n, err = this.Conn.Write(data)

	if n != int(pkgLen) || err != nil {

		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	return
}
