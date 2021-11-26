package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("获取：", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//一旦main函数退出，就释放网络连接
	defer response.Body.Close()
	//读取响应中的所有数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(body))
	fmt.Println(len(body)) //字节切片的大小与页面的大小相同
}

func main() {
	go responseSize("https://www.cnblogs.com/")
	go responseSize("https://www.infoq.cn/")
	go responseSize("https://www.baidu.com/")
	time.Sleep(5 * time.Second)

	// var myChannel chan float64
	// myChannel = make(chan float64)

	// myChannel := make(chan float64)
	// myChannel<-3.14

	// <- myChannel

}
