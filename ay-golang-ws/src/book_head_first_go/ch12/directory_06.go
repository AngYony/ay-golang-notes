package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//获取wy目录下包含的每个文件或子目录（仅子目录，非嵌套目录）
	files, err := ioutil.ReadDir("wy")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//判断这个文件是否是目录
		if file.IsDir() {
			fmt.Println("目录：", file.Name())
		} else {
			fmt.Println("文件：", file.Name())
		}
	}
}
