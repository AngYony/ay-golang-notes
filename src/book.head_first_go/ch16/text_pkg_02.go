package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "这是一个模板"
	//基于文本创建一个新的template值
	tmpl, err := template.New("test").Parse(text)
	check(err)
	//将模板写入终端
	err = tmpl.Execute(os.Stdout, nil)
	check(err)

}
