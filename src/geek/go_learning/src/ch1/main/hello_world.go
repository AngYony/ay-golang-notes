package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args)>1{
		fmt.Println("hello world",os.Args[1])
	}
	fmt.Println(os.Args)
	
	os.Exit(0)
} 