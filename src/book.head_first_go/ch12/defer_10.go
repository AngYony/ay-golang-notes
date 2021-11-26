package main

import (
	"fmt"
	"math/rand"
)

func awardPrize() {
	//产生一个1到3之间的随机数
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("产生1")
	} else if doorNumber == 2 {
		fmt.Println("2222")
	} else if doorNumber == 3 {
		fmt.Println("3333")
	} else {
		panic("未知错误")
	}
}
