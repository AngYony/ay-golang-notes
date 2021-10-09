package array_test

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int             // 声明数组
	arr1 := [4]int{1, 2, 3, 4} // 声明并初始化
	arr2 := [...]int{3, 4, 5}  // 省略大小，自动计算
	fmt.Println(arr, arr1, arr2)

	for i, i2 := range arr2 {
		fmt.Println(i, i2)
	}
}
