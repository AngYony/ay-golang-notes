# flag

用于操作命令行参数。



```
func main() {
	// 接收用户输入的命令参数，来监听对应的ip和端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	fmt.Sprintf("%s:%d", *IP, *Port)
}

```

使用说明：

```
cd 当前目录
go build main.go
main.exe -h
main.exe -port 50053
```

