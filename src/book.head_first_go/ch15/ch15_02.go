package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	//将要添加到响应的内容转换为byte切片
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello,web")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "aaa web")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "bbb,web")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/aaa", frenchHandler)
	http.HandleFunc("/bbb", hindiHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
