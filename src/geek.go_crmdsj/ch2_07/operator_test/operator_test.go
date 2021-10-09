package operator_test

import (
	"fmt"
	"testing"
)

// 可读、可写、可执行
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {

	fmt.Println(Readable, Writable, Executable)

	a := 7 // 0111
	// 清除读的功能
	a = a & ^Readable
	// 清除写的功能
	a = a & ^Writable

	fmt.Println(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

}
