# gorm 增删改查推荐语句

或参阅官方文档：[GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/)



## 新增操作

### 单条新增

推荐为Create方法传入类型的指针（地址）。

```go
user:= User{
    Name:"刘备",
}
result:=db.Create(&user)
// 获取自增的ID
fmt.Println(user.ID)
// 获取受影响行数
fmt.Println(result.RowsAffected)
```

### 批量新增





## 更新操作

### Update()

Update()方法一般用于更新单个列的值。

```go
db.Model(&User{ID:1}).Update("Name","张三")
```

### Updates()

Updates()方法一般用于更新多个列的值。

```go
db.Model(&User{ID:1}).Updates(User{Name:"张飞", Age:30 })
```

==注意：Updates方法不会直接更新零值。== 有两种解决方式。

#### 更新为零值

**方式一：将struct的变量的类型设置为指针类型。**

例如：

```go
type User struct {
    Email  *string
}

// 将Email列更新为零值（空字符串）
empty:=""
db.Model(&User{ID:1}).Updates(User{Email:&empty})
```

**方式二：将struct的变量的类型声明为sql.NULLXXX**

例如：

```go
type Product struct {
    Remark sql.NullString
}

db.Model(&product).Updates(Product{Price:200,Remark:sql.NullString{String:"", Valid: true}})
```

