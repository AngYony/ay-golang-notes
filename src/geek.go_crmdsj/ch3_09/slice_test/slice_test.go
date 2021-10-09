package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

func TestSlicedShareMemory(t *testing.T) {
	ch := []string{"A", "B", "C", "D", "E", "F", "G"}
	c2 := ch[3:6]
	fmt.Println(c2, len(c2), cap(c2))
	// 修改c2的值
	c2[0] = "wy"
	fmt.Println(c2)
	// 源数组也跟着修改了
	fmt.Println(ch)

}
