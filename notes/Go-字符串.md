# Go - 字符串

其他编程语言中的字符串由字符组成，==而Go语言中的字符串是由字节组成的==。

一个字符串是一个不可改变的字节序列，它是一个只读的字节数组，因此使用len()函数时，获取的是字符串字节数，而不是字符的个数。

Go中字符串的关键点描述：

- 字符串是一个只读的字节数组，它的每一个元素都不可修改。（虽然是字节数组，长度也是固定的，但字符串长度并不是字符串类型的一部分）
- for range等语法并不能支持非UTF8编码的字符串的遍历。
- 字符串不能赋值为nil，也不能跟nil比较。

Go语言字符串的底层结构在reflect.StringHeader中定义：

```go
type StringHeader struct {
	Data uintptr
	Len int
}
```

字符串其实是一个结构体，因此字符串的赋值操作也就是reflect.StringHeader结构体的复制过程，并不会涉及底层字节数组的复制。

字符串虽然不是切片，但是支持切片操作，不同位置的切片底层访问的是同一块内存数据（因为字符串是只读的，所以相同的字符串面值常量通常对应同一个字符串常量）。



## 字节数

在Go语言中，内置的 len 函数返回的是字符串的**字节数**，而不是文字符号的数目。而下标访问操作 s[i] 则取得第i个字符，其中 0<=i<len(s)。

注意：字符串的第 i 个字节不一定就是第 i 个字符，因为非 ASCII 字符的 UTF-8 码点需要两个字节或多个字节。

如果需要字符串的字符长度，则应该使用unicode/utf8包的RuneCountInString函数。此函数将返回正确的字符数，而不考虑用于存储每个字符的字节数。

```go
asciiString := "abcde"
utf8String := "中国人"
//输出字节长度
fmt.Println(len(asciiString))	//输出：5
fmt.Println(len(utf8String))	//输出：9
// 获取字符长度
fmt.Println(utf8.RuneCountInString(asciiString)) //输出：5
fmt.Println(utf8.RuneCountInString(utf8String))	 //输出：3
```



## 字符串操作

`s[i:j]`：按照原字符串的**字节**的下标方式来产生新字符串，下标从`i`（含边界值）开始，直到`j`（不含边界值），结果的大小为`j-i`个字节。在使用时，如果省略了i或者j，则取其默认值，操作数i的默认值为字符串的起始位置，值为0；操作数j的默认值为字符串的终止位置，值为len(s)。

### 字符串遍历

#### byte方式（不推荐）

```go
wy := "中国"
for i := 0; i < len(wy); i++ {
	fmt.Printf("%d:[%c] ", i, wy[i])
}
```

输出：

```
0:[ä] 1:[¸] 2:[­] 3:[å] 4:[ 5:[½] 
```

#### rune方式

```go
for index, str := range wy {
	fmt.Printf("%d:%c", index, str)
}
```

输出：

```
0:中3:国
```

推荐使用for...range..对字符串进行遍历，可以正确识别每个字符，包括汉字等复杂字符。







## 字符串字面量

字符串字面量指的是形式上带双引号的字节序列，简单点理解就是一个带双引号的字符串。

Go的源文件总是按UTF-8编码，Go的字符串也会按UTF-8解读，因此在源码中，可以直接将Unicode码点写入到字符串字面量中。

字符串字面量中的常用转义符：

```
\a      响铃
\b      退格
\f      换页
\n      换行
\r      回车
\t      制表符
\v      垂直制表符
\'      单引号（只用在 '\'' 形式的rune符号面值中）
\"      双引号（只用在 "..." 形式的字符串面值中）
\\      反斜杠
```

原生的字符串字面量的书写形式是\`...\`，使用反引号而不是双引号。

在原生的字符串字面量中，转义符不起作用。并且人为的回车后，在处理时会删除回车符（换行符会保留），因此可以将字符串字面量展开多行显示。



## Unicode

Unicode囊括了世界上所有文书体系的全部字符。每一个字符都对应一个叫Unicode码点的标准数字。在Go的术语中，这些字符记号称为文字符号（rune）。

源代码中的文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

Go使用rune类型的值来表示Unicode值。通常，一个符文代表一个字符。（当然也有例外）

Go支持将字符串转换为rune值的切片，并将符文切片转换回字符串。要使用部分字符串，应该将它们转换为rune值的切片，而不是byte值的切片。这样，你就不会意外地抓取符文的部分字节。

```go
asciiString := "abcde"
utf8String := "中国人"
//将字符串转换为rune切片
asciiRunes := []rune(asciiString)
utf8Runes := []rune(utf8String)
//获取每个切片的前几个字符
asciiPartial := asciiRunes[3:]
utf8Partial := utf8Runes[2:]
fmt.Println(string(asciiPartial)) //输出：de
fmt.Println(string(utf8Partial))  //输出：人
```

### UTF-8

UTF-8是一种高效的可变长度的编码方式，它可以用8个、16个或者32个二进制位为单个代码点编码。在可变长度编码方式的基础上，UTF-8沿用了ASCII字符的编码，从而使得ASCII字符可以直接转换为相应的UTF-8编码字符。

 utf8 包提供了两个函数，其中 RuneCountInString 函数能够以符文而不是以字节为单位返回字符串的长度，而 DecodeRuneInString 函数则能够解码字符串的首个字符并返回解码后的符文占用的字节数量。 

```go
zg := "中国"
fmt.Println("字节长度：", len(zg))                    //输出：6
fmt.Println("字符长度：", utf8.RuneCountInString(zg)) //输出：2
c, size := utf8.DecodeRuneInString(zg)
fmt.Printf("第一个字符：%c，其字节数为：%v", c, size) //输出：中和3
```

Go语言提供的关键字range不仅可以迭代元素，而且可以直接解码utf-8编码的字符串。

```go
for i, c := range zg {
    //变量c被赋值为该索引上的代码点（rune）
	fmt.Printf("%v %c\n", i, c) 
}
```





## rune 和 byte

统一码联盟（Unicode Consortium）把名为代码点的一系列数值赋值给了上百万个独一无二的字符。例如，大写字母A的代码点为65，而笑脸表情<img src="https://cdn.ptpress.cn/pubcloud/5B0A982E/ushu/UBb60129159591/online/FBOLb64082fe01ed/Images/31.png" style="width: 2%" width="2%">的代码点则为128515。

Go语言提供了rune（符文）类型用于表示单个统一码代码点，该类型是int32类型的别名。rune相当于go的char。

Go语言还提供了uint8类型的别名byte，这种类型既可以表示二进制数据，又可以表示ASCII定义的英文字符。（ASCII包含128个字符，它是统一码的子集）

```go
A1 := 'A'
var B rune = 'B'
//获取单个字符的代码点
fmt.Printf("%T %[1]d  \n", A1) //输出：int32 65
fmt.Printf("%T \n", B)         //输出：int32，因为rune是int32类型的别名
//根据代码点创建变量
var A2 rune = 65
var pi rune = 960 //π的代码点
//获取对应表示的字符
fmt.Printf("%c %c", A2, pi) //输出：A π
```

代码点值对应rune和int32类型（相当于char）。

综合示例：

```go
func main() {
	s := "abc中国"
	fmt.Println("按照rune(int32)输出：")
	for _, c := range s {
		//按照int32输出结果
		fmt.Printf("%v ", c)
	}
	fmt.Println()

	//将字符串转换为字节数组，并打印每个字节的值
	fmt.Println("按照byte(uint8)类型输出：")
	for _, b := range []byte(s) {
		//按照uint8输出
		fmt.Printf("%v ", b)
	}
	fmt.Println()
	fmt.Println("按照十六进制输出:")
	for _, x := range []byte(s) {
		//按照十六进制输出
		fmt.Printf("%X ", x)
	}
	fmt.Println()
	fmt.Println("输出rune对应的字符：")
	//输出字符
	for _, r := range s {
		fmt.Printf("%c", r)
	}
}
```

输出内容：

```
按照rune(int32)输出：
97 98 99 20013 22269 
按照byte(uint8)类型输出：
97 98 99 228 184 173 229 155 189 
按照十六进制输出:
61 62 63 E4 B8 AD E5 9B BD 
输出rune对应的字符：
abc中国
```

使用`utf8.RuneCountInString`获取字符数量：

```go
zg := "中国"
fmt.Println(utf8.RuneCountInString(zg)) //输出：2
```



## 其他类型与字符串类型之间的相互转换

其他类型转换为string类型，有以下几种方式：

- fmt.Sprintf("%参数",表达式)（推荐）

- 使用strconv包的函数

  ```go
  func FormatBool(b bool) string
  func FormatInt(i int64, base int) string
  func FormatUint(i uint64, base int) string
  func FormatFloat(f float64, fmt byte, prec, bitSize int) string
  func Itoa(i int) string
  ```

- string()

### `[]byte` <=> string

[]byte 转换成string：

```go
string([]byte{72,101})
```

string转换为[]byte：

```go
[]byte("hello")
```

### `[]rune` <=> string

string转[]rune：

```go
[]rune("世界")
```

因为底层内存结构的差异，所以字符串到[]rune的转换必然会导致重新分配[]rune内存空间，然后依次解码并复制对应的Unicode码点值。

[]rune转string：

```go
string([]rune{'世','界'})
```

同样因为底层内存结构的差异，[]rune到字符串的转换也必然会导致重新构造字符串。

【关于字符串与字节数组和rune数组之间转换的内存影响，见《Go高级编程》的1.3字符串部分】

todo: Head Frist Go语言程序设计，附录B中的“更多关于符文的信息部分”

### rune <=> string

rune转字符串：

```go
var pi rune = 'π'
//不进行转换，将会按照int32类型输出值
fmt.Println(pi) //输出：960
//转换为字符串
fmt.Println(string(pi)) //输出：π
```

string转rune：

go不允许将string转换成单个rune，必须将string转换成[]rune。

### int <=> string

将一个整数转换为其对应的字符串的值，例如11转换为“11”。

```go
countdown := 65
//输出该代码点表示的字符
fmt.Printf("%c \n", countdown) //输出：A
//将整数转换为ASCII字符
str := strconv.Itoa(countdown) //输出：65
fmt.Println(str)
//将数值转换为字符串
str = fmt.Sprintf("%v", countdown) //输出：65
fmt.Println(str)
```

将数值字符串转换为数值：

```go
//将ASCII字符转换为整数
c, err := strconv.Atoi("10")
if err != nil {
	fmt.Println("转换错误")
}
fmt.Println(c) //输出：10
```

### bool <=> string（不支持相互转换）

在Go语言中，布尔值并没有与之相等的数字值或者字符串值，因此尝试使用<code>string(false)</code>、<code>int(false)</code>这样的方法来转换布尔值，或者尝试使用<code>bool(1)</code>、<code>bool("yes")</code>等方法来获取布尔值，Go编译器都会报告错误。







## 转义字符

| 字符 | 描述                           |
| ---- | ------------------------------ |
| `\t` | 制表位                         |
| `\n` | 换行                           |
| `\"` | 转义双引号，`\`本身作为转义    |
| `\r` | 回车，不换行，会替换掉顶头字符 |

例如：

```
fmt.Println("张三李\r四")
```

将会输出：

```
四三李
```

这是因为`\r`只回车，不换行，它将从当前行的最前面开始输出，覆盖掉以前内容。



## 字符串常用函数

主要使用到两个包：

- strconv：字符串相关的类型转换
- strings：字符串操作的

| 函数名                              | 描述                                                         |
| ----------------------------------- | ------------------------------------------------------------ |
| len(str)                            | 内置函数，按字节获取字符串的长度，如果按照字符，需要转换为[]rune。 |
| strconv.Itoa(num)                   | 将int类型转换为字符串类型,等价于 strconv.FormatInt(num,10)。 |
| strconv.Atoi(str)                   | 将字符串转换为int，等同于strconv.ParseInt(str,10,0)          |
| []byte(str)                         | 字符串转bype数组                                             |
| string([]byte{...})                 | []byte转字符串                                               |
| strconv.FormatInt(num,b)            | 10进制数转2/8/16进制，第二个参数指定进制位格式               |
| strings.Contains(str1,str2)         | 判断字符串str1里是否包含子串str2                             |
| strings.Count(str1,str2)            | 统计一个字符串str1中有几个子串str2                           |
| strings.EqualFold(str1,str2)        | 比较两个字符串是否相同，不区分大小写，如果区分大小写，可以使用`==` |
| strings.Index(str1,str2)            | 返回子串str2第一次出现在字符串str1中的下标位置，从0开始，没有返回-1，表不存在 |
| strings.LastIndex(str1,str2)        | 返回子串最后一次出现的下标位置，没有返回-1                   |
| strings.Replace(str1,str2,str3,num) | 将字符串str1中的子串str2替换成另一个子串str3，num指定替换的个数，如果num=-1表示全部替换 |
| strings.Split(str,c)                | 将字符串str按照字符c进行分割，返回字符串数组                 |
| strings.ToUpper(str)                | 转换为大写                                                   |
| strings.ToLower(str)                | 转换为小写                                                   |
| strings.TrimSpace(str)              | 去除字符串左右两边空格                                       |
| strings.Trim(str1,str2)             | 去除字符串str1左右两边的指定字符串str2                       |
| strings.TrimLeft(str1,str2)         | 去除字符串左边的字符串                                       |
| strings.TrimRight(str1,str2)        | 去除字符串右边的字符串                                       |
| strings.HasPrefix(str1,str2)        | 判断字符串是否以指定的字符串开头                             |
| strings.HasSuffix(str1,str2)        | 判断字符串是否以指定的字符串结尾                             |

