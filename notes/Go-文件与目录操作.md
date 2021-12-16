# Go-文件与目录操作

文件与目录的操作需要引入os包。os.File封装了文件相关的所有操作。

[IO操作 · Go语言中文文档 (topgoer.com)](https://www.topgoer.com/常用标准库/IO操作.html)

## 文件操作

文件在程序中是以流的形式来操作的。

流：数据在文件和内存之间经历的路径。

输出流：从内存写入到文件中。

输入流：从文件读取到内存中。



### os.OpenFile()

所有的文件操作，都需要先打开文件，获取文件句柄。

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
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件，一般结合O_WRONLY实现内容覆盖
)
```

第三个参数FileMode类型表示文件权限，其格式与你在Unix的ls命令中看到的格式类似。

```go
fmt.Println(os.FileMode(0700)) //输出：-rwx------
```



### 创建文件

方式一，使用os.OpenFile()函数，通过指定选项来创建文件。

```go
file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return
    }
```

方式二，直接调用os.Create()方法来创建文件。

```go
 // 新建文件
    file, err := os.Create("./xxx.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
```

### 写入文件

方式一（推荐）：使用bufio包写入文件。

```go
func wr() {
    // 参数2：打开模式，所有模式d都在上面
    // 参数3是权限控制
    // w写 r读 x执行   w  2   r  4   x  1
    // 如果文件不存在就创建文件
    file, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return
    }
    defer file.Close()
    // 获取writer对象
    writer := bufio.NewWriter(file)
    for i := 0; i < 10; i++ {
        writer.WriteString("hello\n")
    }
    // 刷新缓冲区，强制写出
    writer.Flush()
}
```

方式二，使用file.Write()写入内容。

```go
func main() {
    // 新建文件
    file, err := os.Create("./xxx.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    for i := 0; i < 5; i++ {
        file.WriteString("ab\n")
        file.Write([]byte("cd\n"))
    }
}
```



### 读取文件

方式一（推荐），使用bufio包读取内容。

```go
func re() {
    file, err := os.Open("./xxx.txt")
    if err != nil {
        return
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        if err != nil {
            return
        }
        fmt.Println(string(line))
    }

}
```

也可以使用如下方式读取：

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

方式二，使用file.Read()读取内容。

```go
func main() {
    // 打开文件
    file, err := os.Open("./xxx.txt")
    if err != nil {
        fmt.Println("open file err :", err)
        return
    }
    defer file.Close()
    // 定义接收文件读取的字节数组
    var buf [128]byte
    var content []byte
    for {
        n, err := file.Read(buf[:])
        if err == io.EOF {
            // 读取结束
            break
        }
        if err != nil {
            fmt.Println("read file err ", err)
            return
        }
        content = append(content, buf[:n]...)
    }
    fmt.Println(string(content))
}
```

方式三，使用ioutil包一次性读取文件的全部内容，适用于小文件读取。

```go
package main

import (
   "fmt"
   "io/ioutil"
)

func wr() {
   err := ioutil.WriteFile("./yyy.txt", []byte("www.5lmh.com"), 0666)
   if err != nil {
      fmt.Println(err)
      return
   }
}

func re() {
   content, err := ioutil.ReadFile("./yyy.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(content))
}

func main() {
   re()
}
```



### 拷贝文件

方式一，使用io.Copy()方法。

```go
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)

}
```

方式二，自己实现Copy()的核心代码。

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // 打开源文件
    srcFile, err := os.Open("./xxx.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 创建新文件
    dstFile, err2 := os.Create("./abc2.txt")
    if err2 != nil {
        fmt.Println(err2)
        return
    }
    // 缓冲读取
    buf := make([]byte, 1024)
    for {
        // 从源文件读数据
        n, err := srcFile.Read(buf)
        if err == io.EOF {
            fmt.Println("读取完毕")
            break
        }
        if err != nil {
            fmt.Println(err)
            break
        }
        //写出去
        dstFile.Write(buf[:n])
    }
    srcFile.Close()
    dstFile.Close()
}
```



### 文件或目录是否存在

使用os.Stat()方法来判断文件或路径是否存在。

```go
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
    // 判断返回的错误是否是IsNotExist错误
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
```





 





## 目录操作

### 创建目录

### 删除目录











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

