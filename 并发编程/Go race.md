# Go race 

Go 提供了一个检测并发访问共享资源是否有问题的工具： [race detector](https://blog.golang.org/race-detector)，它可以帮助我们自动发现程序有没有 data race 的问题。



在编译（compile）、测试（test）或者运行（run）Go 代码的时候，加上 race 参数，就有可能发现并发问题。

```
go run -race .\main.go
```

执行上述命令后，如果存在问题，则会输出警告信息：

```shell
PS E:\src\geek.go_bfbcszk\s01_mutex> go run -race .\main.go
==================
WARNING: DATA RACE
Read at 0x00c00012c058 by goroutine 8:
  main.main.func1()
      E:/Wy_Work/AngYony/ay-golang-notes/src/geek.go_bfbcszk/s01_mutex/main.go:18 +0x84

Previous write at 0x00c00012c058 by goroutine 7:
  main.main.func1()
      E:/Wy_Work/AngYony/ay-golang-notes/src/geek.go_bfbcszk/s01_mutex/main.go:18 +0x9d

Goroutine 8 (running) created at:
  main.main()
      E:/Wy_Work/AngYony/ay-golang-notes/src/geek.go_bfbcszk/s01_mutex/main.go:14 +0xeb
...
```



运行 `go tool compile -race -S main.go`，可以查看计数器例子的代码：

```
PS E:\Wy_Work\AngYony\ay-golang-notes\src\geek.go_bfbcszk\s01_mutex> go tool compile -race -S .\main.go
"".main STEXT size=426 args=0x0 locals=0x70 funcid=0x0
        0x0000 00000 (.\main.go:8)      TEXT    "".main(SB), ABIInternal, $112-0
        0x0000 00000 (.\main.go:8)      MOVQ    TLS, CX
        0x0009 00009 (.\main.go:8)      PCDATA  $0, $-2
        0x0009 00009 (.\main.go:8)      MOVQ    (CX)(TLS*2), CX
        0x0010 00016 (.\main.go:8)      PCDATA  $0, $-1
...
```





> 编译的时候不能发现data race,但是编译的时候可以开启race参数，这样编译后的程序在运行时就可以data race问题了。
>
> 绝对不要把带race参数编译的程序部署到线上。