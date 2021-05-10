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



## switch

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
switch {
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

