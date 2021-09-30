package initialize

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/hashicorp/consul/api"
)

func InitSrvConn() {
	// 从注册中心获取用户服务的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// https://www.consul.io/api-docs/agent/service#sample-response
	// Service 属性来自于上述链接中的json属性，可以换成其他属性，比如ID
	data, err := client.Agent().ServicesWithFilter(
		fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))

	if err != nil {
		panic(err)
	}
	// 只获取一次值
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}

	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	zap.S().Info(fmt.Sprintf("连接成功！客户端：%s:%d", userSrvHost, userSrvPort))

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", "msg", err.Error())
	}

	userSrvClient := proto.NewUserClient(userConn)

	// todo:一个连接多个groutine共用存在的问题：性能，解法方式：连接池
	global.UserSrvClient = userSrvClient

}
