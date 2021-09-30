package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	//consul服务所在的地址信息，安装的服务器所在的地址
	cfg.Address = "192.168.171.223:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	// 和new()效果一样，都是实例化对象
	check := &api.AgentServiceCheck{
		// 检查的链接，consul服务能够访问的本地计算机地址，注意：不是127.0.0.1
		HTTP:                           "http://10.112.51.198:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s", // 每 5S 检查一次
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

// AllServices 获取全部服务
func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.171.223:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)

	}

}

// 过滤获取指定的服务
func FilterServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.171.223:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// https://www.consul.io/api-docs/agent/service#sample-response
	// Service 属性来自于上述链接中的json属性，可以换成其他属性，比如ID
	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)

	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)

	}

}

func main() {
	// 虚拟机能够访问的本地电脑的地址
	// _ = Register("10.112.51.198", 8021, "user-web", []string{"mxshop", "wy"}, "user-web")
	// AllServices()
	FilterServices()
}
