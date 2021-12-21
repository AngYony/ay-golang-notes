package process

import (
	"encoding/json"
	"fmt"
	"net"
	"shangguigu.Go_rmdsz/s14_chatroom/common/message"
	"shangguigu.Go_rmdsz/s14_chatroom/server/model"
	"shangguigu.Go_rmdsz/s14_chatroom/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

// 处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	// 从mes 反序列化为loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	// 判断用户名和密码，并将信息序列化返回
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册后再使用..."
	// }

	// redis 中验证数据是否存在
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
	}
	// 将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)

	// 将消息再次序列化发送出去
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail,err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	// 发送回客户端
	err = tf.WritePkg(data)
	return

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 从mes 反序列化为 registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// modelUser := model.User{
	// 	UserId:   registerMes.User.UserId,
	// 	UserPwd:  registerMes.UserPwd,
	// 	UserName: registerMes.UserName,
	// }

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.UserName, "登录成功")
	}

	// 将loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)

	// 将消息再次序列化发送出去
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail,err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	// 发送回客户端
	err = tf.WritePkg(data)
	return
}
