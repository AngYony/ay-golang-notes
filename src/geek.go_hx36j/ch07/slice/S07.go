package main

import "fmt"

func main() {
	//创建长度为5的切片
	s1 := make([]int, 5)
	fmt.Printf("s1切片的长度：%d \n", len(s1)) //输出：5
	fmt.Printf("s1切片的容量：%d \n", cap(s1)) //输出：5
	fmt.Printf("s1切片的值：%d \n", s1)       //输出：[0 0 0 0 0]

	//创建长度为5，容量为8的切片
	s2 := make([]int, 5, 8)
	fmt.Printf("s2切片的长度：%d \n", len(s2)) //输出：5
	fmt.Printf("s2切片的容量：%d \n", cap(s2)) //输出：8
	fmt.Printf("s2切片的值：%d \n", s2)       //输出：[0 0 0 0 0]

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("s4切片的长度：%d \n", len(s4)) //输出：3
	fmt.Printf("s4切片的容量：%d \n", cap(s4)) //输出：5
	fmt.Printf("s4切片的值：%d \n", s4)       //输出：[4 5 6]

}
