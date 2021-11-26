package main

import (
	"fmt"
	"log"
)

func socialize() error {
	defer fmt.Println("滚蛋")

	fmt.Println("Hello")
	return fmt.Errorf("啥玩意")
	fmt.Println("我去")
	return nil
}

func main() {
	err := socialize()
	if err != nil {
		log.Fatal(err)
	}
}
