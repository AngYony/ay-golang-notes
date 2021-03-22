//echo4 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"strings"
)
//变量sep和n是指向标识变量的指针，因此必须通过*sep和*n来访问
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("结束")
	}
}
