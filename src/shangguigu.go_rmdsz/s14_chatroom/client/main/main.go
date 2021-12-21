package main

import (
	"fmt"
	"os"
	"shangguigu.Go_rmdsz/s14_chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	// 接收用户的选项
	var key int
	// 判断是否还继续选择菜单

	for {
		fmt.Println("--------欢迎登录多人聊天系统-------------")
		fmt.Println("1：登录聊天室")
		fmt.Println("2：注册用户")
		fmt.Println("3：退出系统")
		fmt.Println("请选择（1-3）：")

		// 记录用户输入的内容
		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			panic(err)
		}

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的Id")
			fmt.Scanf("%d\n", &userId) // Scanf参数必须加\n
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户Id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名称：")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

	}

	// // 根据用户的输入，显示新的提示信息
	// if key == 1 {
	//
	// 	// fmt.Scanln(&userPwd)
	// 	_ = login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功！")
	// 	// }
	//
	// } else if key == 2 {
	// 	// 用户注册逻辑
	//
	// }

}
