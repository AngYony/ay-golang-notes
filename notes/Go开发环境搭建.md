# Go开发环境搭建



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



## go mod