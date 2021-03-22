package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//打开数据文件进行读取
	file, err := os.Open("D://data.txt")
	//如果打开文件时出现错误，报告错误并退出
	if err != nil {
		log.Fatal(err)
	}
	//为文件创建一个新的扫描器
	scanner := bufio.NewScanner(file)
	//循环到文件结尾，scanner.Scan()会返回false
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
