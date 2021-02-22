# Go语言从入门到实战

## Go语言简介

简单，GO只有25个关键字

高效，编译的强类型语言

## 应用程序入口

1、必须是main包：package main

2、必须是main方法：func main()

3、文件名不一定是main.go

## 退出返回值

与其他主要编程语言的差异：

Go中main函数不支持任何返回值

通过 os.Exit 来返回状态

```go
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello world")
	os.Exit(0)
}
```



## 获取命令行参数

与其他主要编程语言的差异：

- main函数不支持传入参数
- func main(~~arg[] String~~)
- 在程序中直接通过os.Args获取命令行参数

```go
package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args)>1{
		fmt.Println("hello world",os.Args[1])
	}
	fmt.Println(os.Args)
	
	os.Exit(0)
}
```

运行上述程序，输入命令行参数值“wang”，得到结果如下：

```powershell
> go run .\hello_world.go wang
hello world wang
```



## 编写测试程序

1、源码文件以_test结尾：xxx_test.go

2、测试方法名以Test开头：func TestXXX(t * testing.T){...}