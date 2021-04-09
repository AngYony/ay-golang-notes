# Go-文件与目录操作

文件与目录的操作需要引入os包。



## 文件操作

### os.OpenFile()

该函数具有三个参数。第一个参数是一个字符串文件名；第二个参数是文件的操作选项，一个int“标志”；第三个参数表示要设置的文件的权限。

该函数的定义源码如下：

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
```

该函数的调用示例代码：

```go
options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
file, err := os.OpenFile("wy.txt", options, os.FileMode(0600))
```

第二个参数是由几个标志常量按照位操作计算得到的。

```go
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)
```

第三个参数FileMode类型表示文件权限，其格式与你在Unix的ls命令中看到的格式类似。

```go
fmt.Println(os.FileMode(0700)) //输出：-rwx------
```



### 读取文件内容

简单示例：

```go
func main() {
	//打开数据文件进行读取
	file, err := os.Open("D://data.txt")
	//如果打开文件时出现错误，报告错误并退出
	if err != nil {
		log.Fatal(err)
	}
	//为文件创建一个新的扫描器
	scanner := bufio.NewScanner(file)
	//循环到文件结尾，scanner.Scan()会返回false，此处类似于while
	for scanner.Scan() { //从文件中读取一行
		fmt.Println(scanner.Text()) //打印该行
	}

	//关闭文件以释放资源
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//如果扫描文件时出现错误，报告并退出。
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
```



## 目录操作



### 递归读取目录下的所有目录和文件

```go
func reportPanic() {
	//存储panic返回的值
	p := recover()
	//如果没有panic，recover返回nil
	if p == nil {
		return
	}
	//获取底层的error值
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

//递归函数，接收要扫描的路径
func scanDirectory(path string) {
	fmt.Println(path) //打印当前目录

	//获取包含目录内容的切片
	files, err := ioutil.ReadDir(path)
	if err != nil {
		//发生错误，直接引发崩溃
		panic(err)
	}

	for _, file := range files {
		//用斜杠将目录路径和文件名连接起来
		filePath := filepath.Join(path, file.Name())
		//如果是一个目录
		if file.IsDir() {
			//递归调用scanDirectory函数，使用子目录的路径
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	defer reportPanic()
	scanDirectory("wy")
}
```

