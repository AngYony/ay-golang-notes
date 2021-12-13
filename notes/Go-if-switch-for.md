# Go - if - switch - for



## if

和C#中的使用规则一样，只是语句的格式要求严格一些。

需要注意 “else if”语句需要紧跟着上一个“}”的后面，不能换行。并且每一个“{”都必须和条件语句在同一行，也不能换行。

```go
var flag = "A"
if flag == "A" {
	fmt.Println("值为A")
} else if flag == "B" {
	fmt.Println("值为B")
} else {
	fmt.Println("未知值")
}
```

Go语言允许在if语句中使用简短声明变量，该变量可以在if语句的所有else分支中访问。

```go
if num := rand.Intn(3); num == 0 {
	fmt.Println("随机生成值为0")
} else if num == 1 {
	fmt.Println("随机生成值为1")
} else {
	fmt.Println("其他值：", num)
}
```

if的后面必须是条件表达式不能是赋值语句。下述语句在C#中，可以正确运行，但在go语言中，将会编译失败。

```go
// 错误示例，将会编译失败
if b=false {
...
}
```



## switch

Go语言中的switch语句，case后面不需要使用break关键字。

case后面是一个表达式，可以是常量值、变量、一个有返回值的函数、比较表达式等。

### switch常见形式

形式一：

```go
var flag = "A"
switch flag {
case "A":
	fmt.Println("值为A")
case "B":
	fmt.Println("值为B")
default:
	fmt.Println("未知值")
}
```

形式二：每个case分支单独设置比较条件。

```go
switch flag:="C";{      //相当于switch flag:="c"; true{...}
case flag == "A":
	fmt.Println("值为A")
case flag == "B":
	fmt.Println("值为B")
default:
	fmt.Println("未知值")
}
```

形式三：case分支值合并的形式。

```go
switch num := rand.Intn(5); num {
case 1:
	fmt.Println("值为1")
case 2, 3, 4:
	fmt.Println("值为2,3,4中的一个")
default:
	fmt.Println("其他值")
}
```



### switch 语句中的关键字：fallthrough

fallthrough表示穿透，只能穿透一层case。

fallthrough关键字用于执行下一个分支的代码，而不用考虑下一个分支是否满足条件。

```go
switch {
case flag == "A":
	fmt.Println("值为A")
	fallthrough
case flag == "B":
	fmt.Println("值为B")
default:
	fmt.Println("未知值")
}
```

输出结果：

```
值为A
值为B
```



### switch 做类型的推断匹配

```go
func main() {

	var x interface{} = func(x int) string {
		return fmt.Sprintf("d:%d", x)
	}

	switch v := x.(type) {

	case func(int) string:
		fmt.Println(v(100))
	case int:
		fmt.Println("int类型")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("无类型")
	}
}
```





## for

在Go语言中，循环只有for关键字可用。

形式一：for后面直接跟条件。

```go
var count = 10
for count > 0 {
	fmt.Println(count)
	count--
}
```

形式二：省略条件，实现无线循环，通过break跳出循环。

```go
var count = 10
for {
	if count == 0 {
		break
	}
	fmt.Println(count)
	count--
}
```

形式三：与C#中的for一致

```go
for count := 10; count > 0; count-- {
	fmt.Println(count)
}
```

形式四：for..range数据迭代

```
data := [4]string{"a", "b", "c", "d"}
for i, s := range data {
	fmt.Println(i, ":", s)
}
```



### for...range.. 在遍历数组和切片时的区别



```go
data := [3]string{"a", "b", "c"}
// 遍历数组时，s会先复制data中的每个元素数据，所以即使改变了每个元素的值，输出s时，依然为最初元素值
for i, s := range data {
	if i == 0 {
		data[0] += "a"
		data[1] += "b"
		data[2] += "c"
	}
	fmt.Println(i, ":", s, ":", data[i])
}
fmt.Println(data)
// 遍历切片时，一旦改变了切片中的元素的值，每次遍历都会取最新元素的值
for i, s := range data[:] {
	if i == 0 {
		data[0] += "a"
		data[1] += "b"
		data[2] += "c"
	}
	// s是最新的元素的值
	fmt.Println(i, ":", s, ":", data[i])
}
```

输出结果：

```
0 : a : aa   
1 : b : bb   
2 : c : cc   
[aa bb cc]   
0 : aa : aaa 
1 : bbb : bbb
2 : ccc : ccc
```



### break 和 continue、goto

和C#中的作用相同，并支持通过标签指明要break和continue的语句块。

break：默认情况下，break只终止当前层的for循环，可以结合label标签，终止指定层次的循环代码块。

continue：结束本次循环，继续下一次循环。支持标签跳转。

goto：结合标签使用，跳转到指定标签语句，不建议使用goto语句。

```go
b1:
	for x := 0; x < 3; x++ {
	b2:
		for y := 0; y < 3; y++ {
			if x == 2 {
				break b1 // 将会跳出x循环
			}
			for z := 0; z < 3; z++ {
				if z == 2 {
					continue b2 //将会跳过此次y循环，进行下一次y遍历
				}
				fmt.Println(x, "-", y, "-", z)
			}
		}

	}
```



