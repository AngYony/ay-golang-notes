# Struct类型的切片进行排序操作

原理是调用sort包的sort.Sort()方法，该方法接收一个接口参数interface，需要实现该接口。

```go
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// 基于Student类型声明一个切片
type StudentSlice []Student

// 实现sort包中的interface接口

// 返回切片的个数
func (stus StudentSlice) Len() int {
	return len(stus)
}

// 定义排序规则
func (stus StudentSlice) Less(i, j int) bool {
	return stus[i].Age < stus[j].Age // 升序排列
}

// 定义交换方式
func (stus StudentSlice) Swap(i, j int) {
	//temp := stus[i]
	//stus[i] = stus[j]
	//stus[j] = temp
    stus[i],stus[j]=stus[j],stus[i]
}

func main() {
	var stus = StudentSlice{
		{
			Name: "孙悟空",
			Age:  30,
		},
		{
			Name: "猪八戒",
			Age:  20,
		},
		{
			Name: "唐僧",
			Age:  50,
		},
	}

	sort.Sort(stus)
	fmt.Println("排序之后的结果:", stus)

}
```

