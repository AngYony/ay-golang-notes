package main

import (
	"errors"
	"fmt"
)

var (
	//errors.New函数返回的是指针
	ErrBounds = errors.New("什么")
)

func getErr() error {
	return ErrBounds
	//return errors.New("错误")
}

func main() {
	err := getErr()

	switch err {
	case ErrBounds:
		fmt.Println("好")
	default:
		fmt.Println(err)
	}
}
