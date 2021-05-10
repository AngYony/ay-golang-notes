package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4)    //追加一个元素
	mySlice = append(mySlice, 5, 6) //同时追加两个元素
	//fmt.Println(mySlice)            //输出：[1 2 3 4 5 6 7 8]
	// 在切片末尾追加一个切片，切片需要解包
	mySlice = append(mySlice, []int{7, 8}...)
	//fmt.Printf("%#v", mySlice) //输出：[]int{1, 2, 3, 4, 5, 6, 7, 8}

	mySlice = append([]int{0}, mySlice...)          //在开头添加一个元素0
	mySlice = append([]int{-3, -2, -1}, mySlice...) // 在开头添加一个切片
	//fmt.Printf("%#v", mySlice)                      //输出：[]int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	//在索引为2的位置插入66
	mySlice = append(mySlice[:2], append([]int{66}, mySlice[2:]...)...)
	fmt.Println(mySlice)
	//在索引为5个位置插入切片
	mySlice = append(mySlice[:5], append([]int{51, 52, 53}, mySlice[5:]...)...)
	fmt.Println(mySlice)

	//向切片中添加元素
	mySlice = append(mySlice, 0)     //切片扩展一个空间
	copy(mySlice[2+1:], mySlice[2:]) //mySlice[2:]向后移动一个位置
	mySlice[2] = 100                 //设置新添加的元素

	fmt.Println(mySlice)

	//向切片中添加切片
	var x = []int{97, 98, 99}
	mySlice = append(mySlice, x...)       //为新切片扩展足够的空间
	copy(mySlice[2+len(x):], mySlice[2:]) //mySlice[2:]向后移动len(x)个位置
	copy(mySlice[2:], x)                  //复制新添加的切片
	fmt.Println(mySlice)

	//从尾部删除元素
	mySlice = mySlice[:len(mySlice)-1] //删除尾部1个元素
	fmt.Println(mySlice)

	//从开头删除元素
	mySlice = mySlice[1:] //删除开头一个元素
	fmt.Println(mySlice)

	//append()原地完成删除
	mySlice = append(mySlice[:0], mySlice[1:]...) //删除开头一个元素
	fmt.Println(mySlice)

	//使用copy()删除开头的元素
	mySlice = mySlice[:copy(mySlice, mySlice[2:])] //删除开头2个元素
	fmt.Println(mySlice)

	//删除中间的元素
	mySlice = append(mySlice[:2], mySlice[2+1:]...) //从索引为2的位置删除1个元素
	fmt.Println(mySlice)
	mySlice = mySlice[:3+copy(mySlice[3:], mySlice[3+4:])] //从索引为3的位置处删除4个元素
	fmt.Println(mySlice)

}
