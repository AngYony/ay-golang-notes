package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

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
