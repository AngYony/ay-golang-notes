package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

//定义将要传递给模板的类型
type Guestbook struct {
	SignatureCount int      //签名的总数
	Signatures     []string //签名数组
}

func check(err error) {
	if err != nil {
		//发生错误，直接结束程序
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")    //从文件中读取签名，保存到数组中
	html, err := template.ParseFiles("view.html") //基于view.html的内容创建一个模板
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures), //保存签名数组中的数量
		Signatures:     signatures,      //保存签名本身
	}
	//将Guestbook struct数据插入模板，并将结果写入responseWriter中
	err = html.Execute(writer, guestbook)
	check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	//基于new.html文件的内容创建一个模板
	html, err := template.ParseFiles("new.html")
	check(err)
	//将模板写入responseWriter，这里传入nil，表示没有要插入的数据
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	//获取HTML页面中的表单字段的值
	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE

	//打开文件进行写入，如果文件存在，就进行追加，如果不存在，就创建它
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)

	//将表单字段内容添加到文件中
	_, err = fmt.Fprintln(file, signature)
	check(err)

	err = file.Close()
	check(err)

	//将浏览器页面重定向到其他页面
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	//保存文件的每一行字符串
	var lines []string

	file, err := os.Open(fileName)
	//如果得到一个错误，表明文件不存在
	if os.IsNotExist(err) {
		return nil
	}
	check(err)

	defer file.Close()
	//为文件内容创建一个扫描器
	scanner := bufio.NewScanner(file)
	//读取文件中的每一行
	for scanner.Scan() {
		//将其文本附加到切片
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines

}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
