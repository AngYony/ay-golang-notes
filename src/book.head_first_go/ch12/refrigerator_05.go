package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			//如果在切片中找到字符串，返回true
			return true
		}
	}
	//如果找不到字符串，返回false
	return false
}

//Refrigerator类型基于字符串切片
type Refrigerator []string

//模拟打开冰箱
func (r Refrigerator) Open() {
	fmt.Println("正在打开冰箱")
}
func (r Refrigerator) Close() {
	fmt.Println("正在关闭冰箱")
}
func (r Refrigerator) FindFood(food string) error {
	r.Open()
	if find(food, r) {
		fmt.Println("找到：", food)
	} else {
		defer r.Close()
		return fmt.Errorf("%s 未找到", food)
	}
	defer r.Close()
	return nil
}

func main() {
	fridge := Refrigerator{"牛奶", "面包", "油条"}
	for _, food := range []string{"牛奶", "油2条"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
