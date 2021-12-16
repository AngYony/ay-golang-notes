# flag

用于操作命令行参数，替代os.Args。

```go
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

综合示例：

```go
func main() {
	var user string
	var age int
	flag.StringVar(&user, "u", "", "用户名，默认为空字符串")
	flag.IntVar(&age, "a", 10, "年龄，默认为10")
	// 必须调用该方法进行转换，才能绑定输入的值
	flag.Parse()
	fmt.Println(user, age)
}
```

打开终端，编译该文件，并执行。

```
go build -o flag.exe main.go
./flag.exe -u wy -a 111
输出：
wy 111
```

