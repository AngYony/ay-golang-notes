package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// var flag = "A"
	// if flag == "A" {
	// 	fmt.Println("值为A")
	// } else if flag == "B" {
	// 	fmt.Println("值为B")
	// } else {
	// 	fmt.Println("未知值")
	// }

	// switch flag {
	// case "A":
	// 	fmt.Println("值为A")
	// case "B":
	// 	fmt.Println("值为B")
	// default:
	// 	fmt.Println("未知值")
	// }

	// switch {
	// case flag == "A":
	// 	fmt.Println("值为A")
	// 	fallthrough
	// case flag == "B":
	// 	fmt.Println("值为B")
	// default:
	// 	fmt.Println("未知值")
	// }
	// var count = 10
	// for {
	// 	if count == 0 {
	// 		break
	// 	}
	// 	fmt.Println(count)
	// 	count--
	// }
	// for count := 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }

	// if num := rand.Intn(3); num == 0 {
	// 	fmt.Println("随机生成值为0")
	// } else if num == 1 {
	// 	fmt.Println("随机生成值为1")
	// } else {
	// 	fmt.Println("其他值：", num)
	// }

	switch num := rand.Intn(5); num {
	case 1:
		fmt.Println("值为1")
	case 2, 3, 4:
		fmt.Println("值为2,3,4中的一个")
	default:
		fmt.Println("其他值")
	}
}
