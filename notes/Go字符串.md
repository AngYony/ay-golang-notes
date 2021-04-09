# Go字符串

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



## UTF-8

## 字符串和字节 slice





## 其他类型与字符串类型之间的相互转换

### `[]byte` <=> string

[]byte 转换成string：

```go
string([]byte{72,101})
```

string转换为[]byte：

```go
[]byte("hello")
```







todo: Head Frist Go语言程序设计，附录B中的“更多关于符文的信息部分”