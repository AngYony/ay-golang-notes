# Go - module

由go命令统一的管理，用户不必关心目录结构。

初始化：go mod init

增加依赖：go get

更新依赖：go get [@版本号

]

清除多余依赖：go mod tidy

将旧项目迁移到 go mod：`go mod init` ,  `go build ./...`





`go build ./...`：当前目录和当前目录的所有子目录中go文件全部重新编译。





