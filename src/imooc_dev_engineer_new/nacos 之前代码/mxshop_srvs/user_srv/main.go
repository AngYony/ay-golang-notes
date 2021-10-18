package main

import (
	"flag"
	"fmt"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"
	"mxshop_srvs/user_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	"google.golang.org/grpc/health"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// 使用说明：
/*
cd 当前目录
go build main.go
main.exe -h
main.exe -port 50053

*/

func main() {
	// 接收用户输入的命令参数，来监听对应的ip和端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	// Port := flag.Int("port", 50051, "端口号")
	Port := flag.Int("port", 0, "端口号")
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("ip:", *IP)
	zap.S().Info("port:", *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	// consul服务所在的地址信息，安装的服务器所在的地址
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	// 和new()效果一样，都是实例化对象
	check := &api.AgentServiceCheck{
		// 检查的链接，consul服务能够访问的本地计算机地址，注意：不是127.0.0.1
		GRPC:                           fmt.Sprintf("10.112.51.198:%d", *Port),
		Timeout:                        "5s",
		Interval:                       "5s", // 每 5S 检查一次
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())

	registration.ID = serviceID // global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"ay", "wy"}
	registration.Address = "10.112.51.198"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")

}
