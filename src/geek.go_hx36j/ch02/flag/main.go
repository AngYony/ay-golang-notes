package main

import (
	"flag"
	"fmt"
)

var name string

func init() {

	//第 1 个参数是用于存储该命令参数值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示。
	//第 2 个参数是为了指定该命令参数的名称，这里是name。
	//第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
	//至于第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到。
	flag.StringVar(&name, "name", "everyone", "用户名称.")

}

func main() {
	//或者：直接返回一个已经分配好的用于存储命令参数值的地址
	name2 := flag.String("name2", "默认名称", "用户名称.")

	//用于真正解析命令参数，并把它们的值赋给相应的变量。
	//最好放在main函数的函数体的第一行
	flag.Parse()

	//注意：必须在flag.Parse()函数之后，使用对应的变量，否则无法输出参数中name值
	fmt.Printf("Hello,%s!", *name2)
	fmt.Printf("Hello,%s!", name)

}
