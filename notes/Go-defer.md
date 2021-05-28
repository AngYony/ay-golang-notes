# defer

defer用于延迟函数的调用。

可以将defer关键字放在任何普通函数或方法调用之前，Go将延迟（也就是推迟）执行函数调用，直到当前函数退出之后。

多用于下述场景：

- 始终关闭打开的文件（Open / Close）
- 清理代码，即使在发生错误时也需要运行
- Lock / Unlock
- PrintHeader / PrintFooter

示例：

```go
func socialize() {
	fmt.Println("拜拜")
	defer fmt.Println("再见") //该调用被推迟到socialize退出之后
	fmt.Println("滚蛋")
}
func main() {
	socialize()
}
```

如果在调用fmt.Println（"Goodbye！"）之前添加defer关键字，则在Socialize函数中的所有剩余代码运行之前以及Socialize退出之前，该调用不会运行。

因此输出：

```
拜拜
滚蛋
再见
```

注意：如果一个函数有return语句，那么defer关键字必须出现在return语句之前，才能确保函数调用发生。同时defer只能延迟函数和方法调用，不能延迟其他语句。

<code>defer</code>语句在函数包含多个<code>return</code>语句时特别有用。

**多个defer语句，遵循后进先出的规则。**

例如：

```go
func run() {
	defer fmt.Println(1)
	defer fmt.Println(2)
}
func main() {
	run()
}
```

上述执行后，会先输出2，再输出1，遵循先进后出的规则。

**参数在defer语句时计算。**

例如：

```go
func run() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 3 {
			panic("引发崩溃")
		}
	}
}
```

输出如下内容：

```
3
2
1
0
panic: 引发崩溃
```







