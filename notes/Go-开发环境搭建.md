# Go开发环境搭建

建议：先安装多版本管理工具GVM，再通过GVM安装Go。

必须要配置的环境编辑：

- GOROOT
- GOPATH
- PATH
- GOPROXY



## 配置 GO 多版本管理工具

推荐使用GVM

[voidint/g: Golang Version Manager (github.com)](https://github.com/voidint/g)









## GOROOT、GOPATH 和 GOBIN

- GOROOT：Go 语言安装根目录的路径，也就是 GO 语言的安装路径。
- GOPATH：若干工作区目录的路径。是我们自己定义的工作空间。
- GOBIN：GO 程序生成的可执行文件（executable file）的路径。



可以把 GOPATH 简单理解成 Go 语言的工作目录，它的值是一个目录的路径，也可以是多个目录路径，每个目录都代表 Go 语言的一个工作区（workspace）。

我们需要利于这些工作区，去放置 Go 语言的源码文件（source file），以及安装（install）后的归档文件（archive file，也就是以“.a”为扩展名的文件）和可执行文件（executable file）。

具体来说，如果产生了归档文件（以“.a”为扩展名的文件），就会放进该工作区的 pkg 子目录；如果产生了可执行文件，就可能会放进该工作区的 bin 子目录。

归档文件在Linux下就是扩展名是.a的文件，也就是archive文件。写过C程序的朋友都知道，这是程序编译后生成的静态库文件。



## 配置环境变量：$GOPATH

在1.8版本前，必须设置`$GOPATH`环境变量；1.8版本后（含1.8）如果没有设置将使用默认值。

默认位置：在Unix上，默认为`$HOME/go`，在windows上默认为`%USERPROFILE%/go`，在Mac上，GOPATH可以通过修改~/.bash_profile来设置。

注意：GOPATH不是Go的安装目录。

GOPATH目录也被称为GO工作区目录。

$GOPATH目录约定有三个子目录：

- src：存放源代码
- pkg：编译后生成的文件
- bin：编译后生成的可执行文件（为了方便，可以把此目录加入到$PATH变量中）

例如：go get github.com/beego/bee

对应：$GOPATH/src/github.com/beego/bee



## 配置goproxy代理

推荐：[七牛云 - Goproxy.cn](https://goproxy.cn/)



参考链接：https://goproxy.io/zh/docs/getting-started.html

Windows：

```text
1. 右键 我的电脑 -> 属性 -> 高级系统设置 -> 环境变量
2. 在 “[你的用户名]的用户变量” 中点击 ”新建“ 按钮
3. 在 “变量名” 输入框并新增 “GOPROXY”
4. 在对应的 “变量值” 输入框中新增 “https://goproxy.io,direct”
5. 最后点击 “确定” 按钮保存设置
```



## 配置GO111MODULE的值

可以输入`go env`命令，查看当前Go的相关的环境变量。

go111module 是否使用模块支持的变量，如果设为off ，代表无模块支持，import的包会从gopath下寻找。如果设为on，代表模块支持，会忽略gopath，在go.mod中寻找依赖。

所以如果go111module = 'off',要将项目放在gopath的路径下，并使用go get 安装需要的第三方模块。

如果 go111module = 'on' ，可以go mod init 初始化go.mod文件，再使用go build，会自动下载需要的第三方模块。

配置语法：

```
go env -w GO111MODULE="on"
```

注意：如果是在GOPATH下创建的项目，可能需要设置go111module = 'off'，如果不想设为off，需要使用go mod。

关于GO111MODULE 的相关介绍：https://learnku.com/go/t/39086



