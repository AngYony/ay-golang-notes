package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

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
