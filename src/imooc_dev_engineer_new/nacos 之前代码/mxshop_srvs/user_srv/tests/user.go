package main

import (
	"context"
	"fmt"
	"mxshop_srvs/user_srv/proto"

	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	// 注意所有的全局变量，不能使用:=的形式赋值，否则将会变成新声明一个新的变量。
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		word, err := userClient.CheckPassWord(context.Background(), &proto.PassWordCheckInfo{
			PassWord:          "admin123",
			EncryptedPassWord: user.PassWord,
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(word.Success)
	}

}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		userRsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("BBB%d", i),
			Mobile:   fmt.Sprintf("2222222%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(userRsp.Id)

	}
}

func main() {
	Init()
	TestGetUserList()
	//TestCreateUser()
	conn.Close()
}
