package main

import (
	"book_head_first_go/ch10/calendar" //导入新包
	"fmt"
)

func main() {

	date := calendar.Date{}
	err := date.SetYear(2018)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err) //报告错误，并停止程序运行
	}

	err = date.SetMonth(10)
	if err != nil {
		fmt.Println(err)
	}

	date.SetDay(13)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)
}
