package main
import (
    "database/sql"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)
type User struct {
    ID           uint
    Name         string
    Email        *string
    Age          uint8
    Birthday     *time.Time
    MemberNumber sql.NullString
    ActivatedAt  sql.NullTime
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

func main() {
    // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
    dsn := "root:root123456@@tcp(59.110.216.174:7306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

    // 设置全局Logger，打印SQL
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold: time.Second,   // 慢 SQL 阈值
            LogLevel:      logger.Info, // logger.Silent, // 日志级别
            IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:      true,         // 是否彩色打印
        },
    )


    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger:newLogger,

    })

    if err != nil{
        panic(err)
    }


    // 完成迁移，根据struct创建对应的表
     _ = db.AutoMigrate(&User{})

     user:= User{
         Name:"刘备",
     }
     result:=db.Create(&user)
     // 获取自增的ID
     fmt.Println(user.ID)
     // 受影响行数
     fmt.Println(result.RowsAffected)


     // db.Create(&User{Name:"wy"})
     //
     // db.Model(&User{ID:1}).Update("Name","张三")
     //
     // db.Model(&User{ID:1}).Updates(User{Name:"张飞", Age:30 })
     //
     // empty:=""
     // db.Model(&User{ID:1}).Updates(User{Email:&empty})





}