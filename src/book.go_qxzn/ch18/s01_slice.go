package main

import "fmt"

func main() {
	strs := []string{"AA", "BB", "CC", "DD", "EE", "FF", "GG"}
	fmt.Println(len(strs), cap(strs)) //输出：7 7
	//未指定容量
	str2 := strs[2:4]
	fmt.Println(str2) //输出：[CC DD]
	//未指定容量时，以底层数组的开始元素到末尾可见元素个数来确定
	fmt.Println(len(str2), cap(str2)) //输出：2  5

	str3 := strs[2:3:6] //输出：[CC]
	fmt.Println(str3)
	fmt.Println(len(str3), cap(str3)) //输出：1  4
}
