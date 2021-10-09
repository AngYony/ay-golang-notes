package type_test

import (
	"fmt"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int = 10
	var b int64
	// b=a // 不能隐式进行类型转换
	b = int64(a)
	t.Log(a, b)

}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = 1+2 // 指针不能直接进行运算
	fmt.Println(a, aPtr)
}
