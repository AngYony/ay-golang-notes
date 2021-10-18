package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"nacos_start/config"
)

func main() {
	// server client
	sc := []constant.ServerConfig{
		{
			IpAddr: "192.168.171.223",
			Port:   8848,
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "0acc94cf-4d1b-4be7-8f90-b651b3049e33", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}

	// 将json字符串转换为json对象
	serverConfig := config.ServerConfig{}
	err = json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)
	if err != nil {
		panic(err)
	}
	// 监听配置变化：ListenConfig
	// 在nacos页面修改了配置文件后，本地内容会自动更新下来
	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-web.json",
	// 	Group:  "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("配置文件发生了变化")
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })
	//
	// time.Sleep(time.Second * 1000)
}
