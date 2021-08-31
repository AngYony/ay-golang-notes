package main

import (
	"googlecrawler/crawler/engine"
	"googlecrawler/crawler/zhengai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
