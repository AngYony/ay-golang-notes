# Go - nil 和零值

单词nil是一个名词，意思是“零”或者“无”。而在Go语言中，nil则是一个零值。指针、切片、映射和接口的零值都是nil。

Go语言中的零值是变量没有做初始化时系统默认设置的值。



## nil

### 指针与 nil

如果一个指针没有明确的指向，尝试解引用一个nil将导致程序崩溃。

```go
var no *int      //定义一个指针类型的指针变量，但没有赋值
fmt.Println(no)  //输出：<nil>
fmt.Println(*no) //报错
```

解法方式，添加不为nil的if判断语句。`if no !=nil {...}`

### 函数与 nil

当变量被声明为函数类型时，它的默认值为nil。

```go
var fn func(a, b int) int
fmt.Println(fn == nil) //输出：true
```

### 切片与 nil

如果切片在声明之后，没有使用复合字面量或者内置的make函数进行初始，那么它的值为nil。

不过，==关键字 range 和 len、append 等内置函数，都可以正常处理值为nil的切片==。

```go
func say(names []string) []string {
	return append(names, "AA", "BB")
}
func main() {
	var my []string
	fmt.Println(my == nil) //输出：true

	//将nil直接传递给形参为切片的函数
	abc := say(nil) //不会报错
	fmt.Println(abc) //正常输出：[AA BB]
}
```

虽然不包含任何元素的空切片，和值为nil的切片并不相等，但它们通常可以替换使用。

### map 与 nil

跟切片的情况一样，如果映射在声明之后没有使用复合字面量或者内置的<code>make</code>函数进行初始化，那么它的值将会是默认的<code>nil</code>。

可以对值为nil的映射执行读取操作，但不能执行写入操作。

```go
var soup map[string]int
fmt.Println(soup == nil) //输出：true
//读取操作，不报错
wy, ok := soup["wy"]
if ok {
	fmt.Println("存在wy")
} else {
	fmt.Println(wy) //输出int类型的零值：0
	fmt.Println("不存在，写入看看")
	//写入操作，报错
	soup["wy"] = 1
}
```

基于上述原因，如果函数只需要对映射执行读取操作，那么向函数传入<code>nil</code>来代替空映射是可行的。

### 接口与 nil

声明为接口类型的变量在未被赋值时的零值为<code>nil</code>。

对一个未被赋值的接口变量来说，它的接口类型和值都是<code>nil</code>，并且变量本身也等于<code>nil</code>。

```go
var v interface{}
fmt.Printf("%T %v %v", v, v, v == nil) //输出：<nil> <nil> true
```

与此相对的是，当接口类型的变量被赋值之后，接口就会在内部指向该变量的类型和值。

```go
var p *int
v = p
fmt.Printf("%T %v %v", v, v, v == nil) //输出：*int <nil> false
//使用`%#v`查看变量的类型和值
fmt.Printf("%#v", v)                   //输出：(*int)(nil)
```

上述代码中，输出指针类型的值时，取的是指针类型的零值nil，此时会出现值为nil的变量不等于nil的情况。

因为Go认定接口类型的变量只有在类型和值都为<code>nil</code>时才等于<code>nil</code>，所以即使接口变量的值仍然为<code>nil</code>，但只要它的类型不为<code>nil</code>，那么该变量就不等于<code>nil</code>。

为了避免在比较接口变量和<code>nil</code>时得到出乎意料的结果，最好的做法是明确地使用<code>nil</code>标识符，而不是指向一个包含<code>nil</code>的变量。

下面的代码是通过不使用指针和<code>nil</code>值来完全避免了<code>nil</code>指针解引用，并且跟单纯的<code>nil</code>值相比，布尔值<code>valid</code>的意图也更为清晰。

```go
type number struct {
	value int
	valid bool //零值为false
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

//实现Stringer接口
func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)

}

func main() {
	n := newNumber(33)
	fmt.Println(n) //执行String()方法，输出value的值：33
	e := number{}
	fmt.Println(e) //输出：not set
}
```



## 零值

各类型对应的零值如下表所示：

| 数据类型  | 零值  | 说明                        |
| --------- | ----- | --------------------------- |
| bool      | false |                             |
| string    | ""    | 空字符串                    |
| map       | nil   | 键值对类型的变量的零值是nil |
| int/float | 0     | 所有的数字类型的零值都是0   |
| 指针      | nil   | 指针类型的变量零值是nil     |
| slice     | nil   |                             |
| map       | nil   |                             |
| interface | nil   |                             |
| func      | nil   | 函数类型的变量的零值为nil   |

