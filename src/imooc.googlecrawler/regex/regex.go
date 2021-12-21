package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is wy@163.com@abc.com
email is wang@def.org
email3 is www@vy.com
email4 is 222@abc.com.cn
`

func main() {

	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// 只匹配找到的第一个
	match := re.FindString(text)
	fmt.Println(match)

	// 匹配全部
	strs := re.FindAllString(text, -1)
	fmt.Println(strs)

	// 提取匹配的内容，返回的是一个二维切片
	strs2 := re.FindAllStringSubmatch(text, -1)

	for _, m := range strs2 {
		fmt.Println(m)
	}

}
