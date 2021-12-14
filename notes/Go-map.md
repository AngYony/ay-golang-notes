# Go - 字典（map）

- map中的key不能是slice、map和function类型，因为这些类型无法用==来判断。
- 



## 映射(map)

map相当于C#中的键值对。可以使用任意类型来作为映射的键，前提是该类型允许进行`==`比较操作。

映射是引用类型。

### 声明map

在Go语言中，使用map关键字来创建映射，后面跟着一对包含键类型的方括号（`[]`）。

#### 方式一：使用声明变量的形式或make函数创建映射

```go
var myMap map[string]float64
```

仅仅声明一个映射变量，不会分配内存，需要调用make函数进行初始化后才能使用。

==注意和数组的区别，数组只要声明就可以使用，而map不可以。==

```go
myMap = make(map[string]float64)
```

或者，直接使用短变量名的形式创建映射：

```go
ranks := make(map[string]int)
```

可以为make函数指定第二个形参，用于为指定数量的键预先分配空间，就像分配切片的容量一样。

```go
mymap := make(map[float64]int, 8)
```

映射的元素赋值和C#中的键值对赋值基本一样，通过键访问来获取或设置对应的值。

注意：映射不需要显式的添加键，直接可以通过下述代码进行键和值的设置。

```go
myMap["aa"] = 1.1
myMap["bb"] = 2.2
```



#### 方式二：映射字面量

对于每一个键/值对，包含一个键、一个冒号和值。多个键/值对之间以逗号分隔。

映射字面量示例如下：

```go
myMap := map[string]float64{"aa": 1.1, "bb": 2.2}
```

如果要创建一个空的映射，只需要花括号的内容为空即可。

```go
myMap2 := map[string]float64{}
```

### map的其他操作

#### 增加和更新键值对

直接使用`map["key"]=value`形式即可。

该操作如果key不存在就增加键值对，否则key存在就修改值为value。

#### 删除键值对

使用内置函数delete来删除映射中的键和值。第一个参数为映射变量，第二个参数为键。

```go
delete(myTestMap, "A")
```

如果要清空全部map，可以使用：

```go
map = make(...)
```

上述操作，可以重新分配map空间，使旧值被gc回收。

#### 查询键值对

```go
//如果仅仅是为了判断值是否存在，可以直接使用“_”空白标识符来忽略value
val, ok = myTestMap["C"]
fmt.Println(ok) //输出：false
if !ok {
    fmt.Println("不存在键为C的值")
}
```

#### 遍历映射

使用for...range循环来遍历映射。

==注意：在不进行特殊处理的情况下，映射的键和值是按照随机顺序展示的。如果需要有序输出，需要借助额外的数组或切片，按照键排序后，再输出==

注意：Go在迭代映射时并不保证键的顺序，因此，同样的映射在进行多次迭代时，可能会产生不同的输出。

同时获取键和值：

```go
for key, val := range myTestMap {
	fmt.Println(key, val)
}
```

仅获取键：

```go
for key := range myTestMap {
	fmt.Println(key)
}
```

仅获取值，需要将键赋给“_”空白标识符：

```go
for _, val := range myTestMap {
	fmt.Println(val)
}
```



### map中的零值（默认值）

映射类型的变量本身的零值是nil。如果声明了一个映射变量但是未赋值，它的值是nil。此时如果尝试增加键或值，将会产生一个错误。因此在增加一个新的键/值对之前，需要使用make或者映射字面量来创建一个映射，并且赋值给映射变量。

映射中值的零值由映射中的值的类型决定。零值可以让你更加安全地修改映射的值，即使在没有赋值给它的情况下。

```go
cnts := make(map[string]int)
cnts["aa"]++
cnts["aa"]++
cnts["bb"]++
fmt.Println(cnts)	//输出：map[aa:2 bb:1]
```

### 如何区分已经赋值的值和零值

在访问映射键的时候，会返回可选的第2个布尔类型的值，如果这个键已经被赋值过，那么返回true，否则返回false。

如下所示：

```go
myTestMap := map[string]int{}
myTestMap["A"] = 1
value, ok := myTestMap["A"]
fmt.Println(value, ok) //输出：1 true
myTestMap["B"] = 0     //直接复制为int类型的零值0
value, ok = myTestMap["B"]
fmt.Println(value, ok)     //输出：0 true
value, ok = myTestMap["C"] //直接访问未被赋过值的键
fmt.Println(value, ok)     //输出：0 false
//如果仅仅是为了判断值是否存在，可以直接使用“_”空白标识符来忽略value
_, ok = myTestMap["C"]
fmt.Println(ok) //输出：false
if !ok {
    fmt.Println("不存在键为C的值")
}
```

