package initialize

import (
	"mxshop-api/user-web/global"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
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

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息：%v", global.ServerConfig)
	// fmt.Println(v.Get("name"))
	// 动态监听文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化：%s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息：%v", global.ServerConfig)
	})

}
