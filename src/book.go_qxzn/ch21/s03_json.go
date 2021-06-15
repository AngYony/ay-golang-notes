package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64 //字段必须以大写字母开头
		Msg       string  `json:"message"`
	}

	curiosity := location{-4.5, 11.45, "你好"}

	//Marshal函数只对结构中被导出的字段实施编码
	bytes, err := json.Marshal(curiosity)
	if err != nil {
		os.Exit(1)
	}

	str := string(bytes)
	fmt.Println(str)
}
