# Go数组

数组保存特定数量的元素，不能增长或收缩。

在Go语言中，数组中的元素索引从0开始。

创建数组时，它所包含的所有值都初始化为数组所保存类型的零值（默认值）。

## 声明数组

### 方式一

语法：

```
var 数组变量名 [个数]类型
```

示例：

```go
var myArray [4]string	//创建一个由4个字符串组成的数组
var myIntArray [5] int  //创建一个由5个整数组成的数组
var dates [3]time.Time  //创建一个由3个Time值组成的数组
```

### 方式二：数组字面量

如果事先知道数组要存储的值，可以使用数组字面量的形式初始化数组。

示例：

```go
var notes [3]string = [3]string{"AA", "BB", "CC"}
fmt.Println(notes[0])
myArray := [3]int{1, 2, 3}	//短变量声明形式
fmt.Print(myArray[2])
```



## 获取或设置数组中的元素的值

直接通过索引下标进行设置和读取。

示例：

```go
myArray[0]="one"		//设置元素的值
myArray[1]="two"		
fmt.Println(myArray[2]) //读取元素的值
```





## 数组相关的其他操作

### 获取数组的长度

使用len()函数返回数组的长度（它包含的元素个数）。

```go
var myArr [3]int = [3]int{1, 2}
fmt.Println(len(myArr))		//输出：3
```

声明数组时，长度设置的是几，len()函数就返回几，而不用管其中的元素是否赋值（没有显式设置值的元素将使用默认值）

### 使用for...range循环遍历数组

处理数组中每个元素的一种更安全的方法是使用特殊的for...range循环。

格式：

```go
for index,value := range myArray{
	//循环体
}
```

格式说明：

- index：保存每个元素索引的变量
- value：保存每个元素值的变量
- myArray：正在处理的数组

示例：

```go
for index, value := range notes {
	fmt.Println(index, value)
}
```

输出：

```
0 AA
1 BB
2 CC
```

如果在遍历的过程中，不想使用其中声明的变量index或value，那么可以使用空白标识符（_）来代替变量名。这将导致Go丢弃该值，而不会产生编译器错误。



## 使用“fmt”输出数组元素

```go
var notes [3]string = [3]string{"AA", "BB", "CC"}
fmt.Println(notes)       //输出：[AA BB CC]
fmt.Printf("%#v", notes) //输出：[3]string{"AA", "BB", "CC"}
```





## 数组中元素的默认值

- int类型的元素默认值为0
- string类型的元素默认值为空字符串。