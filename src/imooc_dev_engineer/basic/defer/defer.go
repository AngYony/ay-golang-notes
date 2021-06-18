package main

import (
	"fmt"
	"os"
)

func run() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 3 {
			panic("引发崩溃")
		}
	}
}
func main() {
	_, err := os.Open("abc.txt")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			fmt.Println("未知错误", err)
		}
	}
}
