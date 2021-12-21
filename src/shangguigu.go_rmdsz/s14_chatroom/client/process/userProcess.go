package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shangguigu.Go_rmdsz/s14_chatroom/common/message"
	"shangguigu.Go_rmdsz/s14_chatroom/server/utils"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {

		fmt.Println("net.Dial err=", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3. 创建一个LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	// 4. 将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. 为Data赋值
	mes.Data = string(data)

	// 6. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 创建一个transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误，err=", err)
	}

	mes, err = tf.ReadPkg() // RegisterResMes
	if err != nil {
		fmt.Println("readpdk(conn) err=", err)
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

// 登录校验
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	// 1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {

		fmt.Println("net.Dial err=", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. 创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. 为Data赋值
	mes.Data = string(data)

	// 6. 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. 发送data数据
	// 先发送长度，在发送内容
	// 获取长度并转换为切片
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte // 因为uint32数字存储需要4个字节，所以这里使用的是4个长度的字节数组
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[0:4])

	if n != 4 || err != nil {

		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	// fmt.Printf("客户端，发送消息长度=%d 字节\n，内容：%s", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data)

	if err != nil {

		fmt.Println("conn.Write(data) fail", err)
		return
	}

	// 这里还需要处理服务器端返回的消息

	tf := &utils.Transfer{
		Conn: conn,
	}
	// 处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpdk(conn) err=", err)
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		// 启动一个协程，用于保持和服务器端的通讯，如果服务器有数据推送给客户端，则接收并显示在客户端的终端
		go serverProcessMes(conn)
		// 显示二级菜单
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}

	return

}
