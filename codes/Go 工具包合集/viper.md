# viper



config-debug.yaml

```yaml
name: 'user-web'
port: 8021
mysql:
  host: '127.0.0.1'
  port: 3306

```

对应的类型：

```go
package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}
```



## 读取YAML文件配置映射到类型

```go
func main() {
	v := viper.New()
	v.SetConfigFile("ch01/config-debug.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := &ServerConfig{}
	if err := v.Unmarshal(serverConfig); err != nil {
		panic(err)
	}
	// fmt.Println(v.Get("name"))

	fmt.Println(*serverConfig)
}
```



## 根据系统环境变量加载不同YAML文件配置

```go
// 读取计算机本地配置的环境变量值，配置完环境变量后，需要重启GoLand才能生效
func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func GetEnvInfo2(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	// 读取 PATH 环境变量值，如果是配置的新的环境变量，需要重启GoLand才能生效
	fmt.Println(GetEnvInfo("PATH"))

	// 通过获取系统环境变量的形式，来动态加载不同的配置文件
	isPro := GetEnvInfo2("MXSHOP_Pro")
	configFileName := "ch02/config-debug.yaml"
	if isPro {
		configFileName = "ch02/config-pro.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := &ServerConfig{}
	if err := v.Unmarshal(serverConfig); err != nil {
		panic(err)
	}
	// fmt.Println(v.Get("name"))

	fmt.Println(*serverConfig)

	// 动态监听文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生了变化：", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(serverConfig)
		fmt.Println(*serverConfig)

	})
	// 防止提前退出
	select {}
}
```



## 动态监听配置文件变化

见上述代码。