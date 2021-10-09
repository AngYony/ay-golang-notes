package map_ext_test

import (
	"fmt"
	"testing"
	"time"
)

func TestMapWithFunValue(t *testing.T) {
	// 声明一个map，key为int，值为func
	// m := map[int]func(op int) int{}

	start := time.Now()
	fmt.Println(time.Since(start).Seconds())

}
