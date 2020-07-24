//测试多个.go文件在同一个包中的调用情况
package main

import (
	"chapter2/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("TTTT! %v\n", tempconv.AbsoluteZeroC) //输出：TTTT! -273.15℃
	fmt.Println(tempconv.CToF(tempconv.BoilingC))    //输出：212℉
}

var a = b + c //3：最后把a初始化为3
var b = f()   //2：接着通过调用f()将b初始化为2
var c = 1     //1：首先初始化c为1
func f() int  { return c + 1 }
