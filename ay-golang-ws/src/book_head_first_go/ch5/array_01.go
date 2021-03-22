package main

import "fmt"

func main() {
	var notes [3]string = [3]string{"AA", "BB", "CC"}
	fmt.Println(notes)         //输出：[AA BB CC]
	fmt.Printf("%#v\n", notes) //输出：[3]string{"AA", "BB", "CC"}

	var myArr [3]int = [3]int{1, 2}
	fmt.Println(len(myArr))

	for index, value := range notes {
		fmt.Println(index, value)
	}
}
