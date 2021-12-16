# go 操作命令





## go 操作命令



### go build

- 如果是普通包，当你执行 go build 之后，它不会产生任何文件。如果你需要在 $GOPATH/pkg 下生成相应的文件，那就得执行 go install。
- 如果是 main 包，当你执行 go build 之后，它就会在当前目录下生成一个可执行文件。如果你需要在 $GOPATH/bin 下生成相应的文件，需要执行 go install，或者使用 go build -o 路径/a.exe。
- 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在 go build 之后加上文件名，例如 go build a.go；go build 命令默认会编译当前目录下的所有 go 文件。
- 你也可以指定编译输出的文件名。例如，我们可以指定 go build -o astaxie.exe，默认情况是你的 package 名 (非 main 包)，或者是第一个源文件的文件名 (main 包)。（注：实际上，package 名在 Go 语言规范 中指代码中 “package” 后使用的名称，此名称可以与文件夹名不同。默认生成的可执行文件名是文件夹名。）
- go build 会忽略目录下以 _ 或 . 开头的 go 文件。
- 使用go build编译后的可执行文件，即使拷贝到没有go开发环境的机器上，仍然可以运行。
- 在编译时，编译器会将程序运行依赖的库文件包含在可执行文件中，所以，生成后的可执行文件变大了很多。

### go install

这个命令在内部实际上分成了两步操作：第一步相当于执行go build，生成结果文件 (可执行文件或者 .a 包)，第二步会把编译好的结果移到 `$GOPATH/pkg` 或者 `$GOPATH/bin`。

### go get

这个命令是用来动态获取远程代码包的，目前支持的有 BitBucket、GitHub、Google Code 和 Launchpad。这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行 go install。



### go test

### go list

### go vet







### go run

用于检测数据访问冲突问题，可以使用如下命令：

```shell
go run -race wy.go
```

go run也是先go build，然后再执行。

go run必须在安装了go开发环境的机器上运行，才能够成功。





### go tool pprof

查看性能成本，CPU使用率等



### gofmt

可以执行gofmt命令，将指定的.go文件格式化输出，如果想要将格式化后的内容写入到要格式化的文件中，需要使用-w选项。

```
gofmt -w main.go
```



