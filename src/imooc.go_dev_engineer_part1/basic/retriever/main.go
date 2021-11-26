package main

import (
	"fmt"
	"time"

	"imooc.dev_engineer/retriever/mock"
	real2 "imooc.dev_engineer/retriever/real"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("https://imooc.com")
}

func post(poster Poster) {
	poster.Post("http://baidu.com",
		map[string]string{
			"name":   "张三",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post("baiducom", map[string]string{"contents": "啦啦啦啦啦啦"})
	return "wwww"
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "这是一个假的消息"}
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Chrome 3.2",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//类型断言
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("不是mock retriever")
	}

}
func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	//获取肚子里的类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
