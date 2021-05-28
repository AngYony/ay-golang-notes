# Go - error

## 存储错误消息的变量

根据惯例，Go程序将使用带有Err前缀的变量来存储错误消息。

```go
var (
	//errors.New函数返回的是指针
	ErrBounds = errors.New("什么")
)

func getErr() error {
	return ErrBounds
	//return errors.New("错误")
}

func main() {
	err := getErr()
	//使用switch比较方法返回的错误变量
	switch err {
	case ErrBounds:
		fmt.Println("好")
	default:
		fmt.Println(err)
	}
}
```



## error 接口

error类型本质上只是一个接口：

```go
type error interface{
	Error() string
}
```

因此，如果自定义的类型中，包含一个返回string的Error方法，那么它就满足error接口，并且它是error的值。

按照惯例，自定义error类型通常使用单词Error作为后缀。

自定义error类型：

```go
//定义一个以string为基础类型的类型
type ComedyError string
//满足error接口
func (c ComedyError) Error() string {
	return string(c)
}

//定义一个基础类型为float64的类型
type OverheatError float64
//实现error接口
func (o OverheatError) Error() string {
	return fmt.Sprintf("值：%0.2f", o)
}
//指定函数返回原生error值
func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
        //var wyerr error
        //wyerr = OverheatError(excess)
        //return wyerr
		return OverheatError(excess) //返回的是error，实现的具体类型
	}
	return nil
}

func main() {
	var err error //声明一个error类型的变量
	err = ComedyError("这是一个错误信息")	//ComedyError满足error接口，所以可以赋值给接口变量
	fmt.Println(err)

	err = checkTemperature(121.322, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
```

> error类型像int或者string一样是一个“预定义标识符”，它不属于任何包。它是“全局块”的一部分，这意味着它在任何地方可用，不用考虑当前包信息。



使用断言对已知的类型错误进行处理：

```go
func main() {
	_, err := os.Open("abc.txt")
	if err != nil {
        //对err进行类型断言
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			fmt.Println("未知错误", err)
		}
	}
}
```

