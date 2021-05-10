package main

import "fmt"

func wyPrint(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	var a = []interface{}{123, "abc"}
	wyPrint(a...)
	wyPrint(a)
}

//在map中查询指定key的value
func find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

func inc() (v int) {
	defer func() { v++ }()
	return 42
}


