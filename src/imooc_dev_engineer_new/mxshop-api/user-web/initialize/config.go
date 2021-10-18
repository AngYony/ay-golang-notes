package initialize

import (
	"encoding/json"
	"fmt"
	"mxshop-api/user-web/global"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {

	// 读取 PATH 环境变量值，如果是配置的新的环境变量，需要重启GoLand才能生效
	// fmt.Println(GetEnvInfo("PATH"))

	// 通过获取系统环境变量的形式，来动态加载不同的配置文件
	isPro := GetEnvInfo("MXSHOP_Pro")

	configFileName := "user-web/config-debug.yaml"
	if isPro {
		configFileName = "user-web/config-pro.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	// fmt.Println(v.Get("name"))
	// 动态监听文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
		zap.S().Infof("配置信息：%v", global.ServerConfig)
	})

	// 从 Nacos 中读取配置信息
	// server client
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host, //  "192.168.171.223",
			Port:   global.NacosConfig.Port,
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
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
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}

	// 将json字符串转换为json对象
	// serverConfig := &global.ServerConfig{}
	err = json.Unmarshal([]byte(content), global.ServerConfig)

	if err != nil {
		zap.S().Fatalf("读取nacos配置失败：%s", err.Error())
	}

	fmt.Println(global.ServerConfig)
}
