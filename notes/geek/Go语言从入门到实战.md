# Go语言从入门到实战

## Go语言简介

简单，GO只有25个关键字

高效，编译的强类型语言

## 开发环境构建

GOPATH：在1.8版本前，必须设置这个环境变量；1.8版本后（含1.8）如果没有设置使用默认值

在Unix上，默认为`$HOME/go`，在windows上默认为`%USERPROFILE%/go`，在Mac上，GOPATH可以通过修改~/.bash_profile来设置。

可以输入下述命令检测版本：

```
go version
```

课程go的版本：1.11.5

```go
package main

import "fmt"

func main(){
	fmt.Println("hello world")
}
```

如果要运行上述代码，需要进入到上述代码文件所在的目录，执行下述命令即可看到输出结果：

```powershell
> go run .\hello_world.go
hello world
```

也可以执行下述命令编译go源码：

```powershell
> go build .\hello_world.go
```

不同的操作系统会生成不同的编译文件，在windows下，会生成.exe文件，可以直接运行该.exe文件，会输入相同的结果：

```powershell
> .\hello_world.exe
hello world
```

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

