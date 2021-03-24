package main

import "fmt"

func main() {
	// var notes []string        //声明一个切片变量
	// notes = make([]string, 7) //创建存储7个字符串的切片
	// notes[0] = "你好"
	// fmt.Printf("%#v\n", notes)

	// fmt.Println(len(notes))

	// wy := []string{"one", "two"}
	// fmt.Println(wy)

	// myArray := [5]int{1, 2, 3, 4, 5}
	// mySlice := myArray[:3]
	// fmt.Println(mySlice) //输出：[1 2 3]
	// myArray[0] = 11
	// fmt.Println(mySlice) //输出：[11 2 3]
	// mySlice[1] = 22
	// fmt.Println(myArray) //输出：[11 22 3 4 5]

	// mySlice := []int{1, 2, 3, 4, 5}
	// mySlice = append(mySlice, 6)    //追加一个元素
	// mySlice = append(mySlice, 7, 8) //同时追加两个元素
	// fmt.Println(mySlice)            //输出：[1 2 3 4 5 6 7 8]

	s1 := []int{1, 2, 3}
	s2 := append(s1, 4, 5)
	s2[0] = 0           //更改了s2的第一个元素，s1没发生变化
	fmt.Println(s1, s2) //此处输出：[1 2 3] [0 2 3 4 5]，s1的值并没有被改变
	s3 := append(s2, 6, 7)
	s4 := append(s3, 8, 9, 10)
	fmt.Println(s3, s4) //此处输出：[0 2 3 4 5 6 7] [0 2 3 4 5 6 7 8 9 10]
	s4[0] = 111         //更改s4的第一个元素，发现对应的s3的第一个元素也发生了改变
	fmt.Println(s3, s4) //输出：[111 2 3 4 5 6 7] [111 2 3 4 5 6 7 8 9 10]

	myFunc(1, "a", "b")

	wy := []string{"AA", "BB"}
	myFunc(1, wy...)
}

func myFunc(a int, b ...string) {
	fmt.Println(a, b)
}
