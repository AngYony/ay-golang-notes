package main

import (
	"fmt"
	"io"
	"net"
	"shangguigu.Go_rmdsz/s14_chatroom/common/message"
	process2 "shangguigu.Go_rmdsz/s14_chatroom/server/process"
	"shangguigu.Go_rmdsz/s14_chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// serverProcessMes 根据客户端发送消息种类不同，决定调用哪个函数处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		up := &process2.UserProcess{
			Conn: this.Conn,
		}

		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}

		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process() error {
	// 循环读取客户端发送的消息
	for {

		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			} else {

				fmt.Println("readpkg err=", err)
				return err
			}
		}

		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
