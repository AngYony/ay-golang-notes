# JSON相关的操作

使用json包用于对象的序列化和反序列化操作。

## 序列化

使用json.Marshal()方法，用于将对象序列化为JSON字符串。

注意：如果一个类中的字段不是大写的，该字段将不会参与序列化操作。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 结构体序列化
	stu := Student{Age: 10, Name: "张三"}
	structData, err := json.Marshal(&stu)
	if err != nil {
		return
	}
	fmt.Println("结构体序列化：", string(structData))

	// map序列化
	m := map[string]interface{}{
		"A": 11,
		"B": "李四",
	}
	mapData, err := json.Marshal(m)
	if err != nil {
		return
	}
	fmt.Println("map序列化：", string(mapData))

	// 切片序列化，切片的每个元素都是一个map
	s := []map[string]interface{}{
		{
			"A1": 10,
			"A2": 20,
		},
		{
			"B1": 30,
			"B2": 40,
		},
		{
			// 值为数组
			"arr": [...]int{1, 2, 3},
		},
	}
	sliceData, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println("切片序列化：", string(sliceData))

}
```

输出：

```
结构体序列化： {"name":"张三","age":10}
map序列化： {"A":11,"B":"李四"}                                   
切片序列化： [{"A1":10,"A2":20},{"B1":30,"B2":40},{"arr":[1,2,3]}]
```



## 反序列化

使用json.UnMarshal()将json字符串反序列化为对象。

```go
// 反序列化为结构体
stuStr := `{"name":"张三","age":10}`
var stu2 Student
err = json.Unmarshal([]byte(stuStr), &stu2)
if err != nil {
	return
}
fmt.Println("反序列化为结构体：", stu2)
// 反序列化为map
mapStr := `{"A":11,"B":"李四"}`
// 反序列化操作不需要make，因为make操作被封装到了Unmarshal中
var m2 map[string]interface{}
err = json.Unmarshal([]byte(mapStr), &m2)
if err != nil {
	return
}
fmt.Println("反序列化map：", m2)
// 反序列化为切片
sliceStr := `[{"A1":10,"A2":20},{"B1":30,"B2":40},{"arr":[1,2,3]}]`
var s2 []map[string]interface{}
err = json.Unmarshal([]byte(sliceStr), &s2)
if err != nil {
	return
}
fmt.Println("反序列化切片：", s2)
```

输出：

```
反序列化为结构体： {张三 10}                                       
反序列化map： map[A:11 B:李四]                                     
反序列化切片： [map[A1:10 A2:20] map[B1:30 B2:40] map[arr:[1 2 3]]]
```





