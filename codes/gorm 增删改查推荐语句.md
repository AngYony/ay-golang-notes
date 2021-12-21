# gorm 增删改查推荐语句

或参阅官方文档：[GORM 指南 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/)

基础模型：

```go
type LetterCombination struct {
	gorm.Model
	LCName     string  `gorm:"column:lc_name;type:char(10)"`
	WordId     string  `gorm:"column:word_id;type:char(36)"`
	WordName   string  `gorm:"column:word_name;type:char(50)"`
	CreatedBy  *string `gorm:"column:created_by;type:varchar(100)"`
	ModifiedBy *string `gorm:"column:modified_by;type:varchar(100)"`
}
```



 

## 新增

可以使用struct或者map进行单条或多条数据的新增操作。

使用map形式更加灵活，并且不会为没有显式指定的列进行值的设定。

### 单条新增

#### struct形式单条新增

```go
lc := LetterCombination{
	WordId:   "ssssss",
	WordName: "bed",
	LCName:   "age",
}
result := db.Create(&lc)
// 获取自增的ID值
fmt.Println(lc.ID)
fmt.Println(fmt.Sprintf("受影响行数：%d", result.RowsAffected)) // 返回插入记录的条数
```

生成的SQL语句：

```mysql
INSERT INTO `letter_combinations` (`created_at`,`updated_at`,`deleted_at`,`lc_name`,`word_id`,`word_name`,`created_by`,`modified_by`) VALUES ('2021-12-04 14:54:32.357','2021-12-04 14:54:32.357',NULL,'age','ssssss','bed',NULL,NULL)
```

#### map形式单条新增

```go
// 使用 Map创建单条记录
db.Model(&LetterCombination{}).Create(map[string]interface{}{
	"LCName": "pro", "WordName": "product",
})
```

生成的SQL语句：

```mysql
INSERT INTO `letter_combinations` (`lc_name`,`word_name`) VALUES ('pro','product')
```

### 批量新增

#### struct形式批量新增

```go
var lcs = []LetterCombination{
	{LCName: "aw", WordName: "haw"},
	{LCName: "ph", WordName: "phone"},
	{LCName: "er", WordName: "docker"},
}
// db.Create(&lcs)
// 数据量大时，使用批次方式进行批量写入
db.CreateInBatches(lcs, 2)
for _, lc := range lcs {
	fmt.Println(fmt.Sprintf("%d:%s", lc.ID, lc.WordName))
}
```

生成的SQL语句：

```mysql
INSERT INTO `letter_combinations` (`created_at`,`updated_at`,`deleted_at`,`lc_name`,`word_id`,`word_name`,`created_by`,`modified_by`) VALUES ('2021-12-04 14:54:32.358','2021-12-04 14:54:32.358',NULL,'aw','','haw',NULL,NULL),('2021-12-04 14:54:32.358','2021-12-04 14:54:32.358',NULL,'ph','','phone',NULL,NULL)
INSERT INTO `letter_combinations` (`created_at`,`updated_at`,`deleted_at`,`lc_name`,`word_id`,`word_name`,`created_by`,`modified_by`) VALUES ('2021-12-04 14:54:32.359','2021-12-04 14:54:32.359',NULL,'er','','docker',NULL,NULL)
```

#### map形式批量新增

```go
db.Model(&LetterCombination{}).Create([]map[string]interface{}{
	{"LCName": "ion", "WordName": "session"},
	{"LCName": "oo", "WordName": "book"},
})
```

生成的SQL语句：

```mysql
INSERT INTO `letter_combinations` (`lc_name`,`word_name`) VALUES ('ion','session'),('oo','book')
```



## 修改

### db.Save()

Save()保存的数据，如果主键不存在就创建一条记录，否则就根据主键修改记录。

```
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
```

Save会根据实体的主键值，来进行新增或者修改操作。

### 更新单列：Update()

Update()方法一般用于更新单个列的值。

```go
db.Model(&User{ID:1}).Update("Name","张三")
```

### 更新多列：Updates()

Updates()方法一般用于更新多个列的值。

#### struct形式

```
// 根据 `struct` 更新属性，只会更新非零值的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
```

==注意：Updates方法不会直接更新零值。== 有两种解决方式。

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



#### map形式

```
// 根据 `map` 更新属性
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
```









## 删除

## 查询

条件过滤支持三种方式：

- struct，最接近面向对象思想，不需要关注数据库表的列，不容易出错，但会忽略零值的处理。
- map，介于struct和string两种方式之间，不会出错，也不会忽略零值的处理。
- string，对应SQL语句的方式，写法最灵活，但是需要正确的指定数据库表的字段。

```go
// 查询一条记录
var user User
db.Where(&User{MyName:"wy"}).First(&user)

//查询多条记录
var users []User
db.Where(&User{MyName:"wy"}).Find(&users)
```

注意：使用struct的形式进行查询操作时，GORM 只会查询非零值字段，这意味着如果您的字段值为 `0`、`''`、`false` 或其他 [零值](https://tour.golang.org/basics/12)，该字段不会被用于构建查询条件，例如：

```go
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu";
```



 