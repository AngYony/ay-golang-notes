# Go文件与目录操作



## 文件操作

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

