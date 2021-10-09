package my_map

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	// 声明并初始化map
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(m1)
	m2 := map[int]int{}
	m2[4] = 15
	fmt.Println(m2)

}
