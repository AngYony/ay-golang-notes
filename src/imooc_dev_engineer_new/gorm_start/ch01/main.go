package main
import (
    "database/sql"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)
type Product struct {
    gorm.Model
    Code  string
    Price uint
    Remark sql.NullString
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
    // _ = db.AutoMigrate(&Product{})

    // 新增一条语句
    db.Create(&Product{Code:"A01",Price:100})

    // 从数据库中读取一条记录
    var product Product
    db.First(&product,1) // 根据主键值进行查找，传递指针为变量修改值
    // 根据其他列查询数据
    db.First(&product,"code = ?","A01") // 查找code值为A01的数据


    // 更新数据
    // 更新单个字段：将product的price更新为200
    db.Model(&product).Update("Price",200)
    // 更新多个字段
    // 方式一：
    db.Model(&product).Updates(Product{Price:200,Code:"A02"}) // 注意：Updates这种写法仅更新非零值字段
    // 方式二：
    db.Model(&product).Updates(map[string]interface{}{"Price":200, "Code":"A03" })

    // 更新字段的值为零值（解决零值写入的问题），struct中字段的类型必须定义为sql.NullString，并使用下述方式赋值
    db.Model(&product).Updates(Product{Price:200,Remark:sql.NullString{String:"", Valid: true}})


    // 删除数据，实质更新的是deleted_at的值，是逻辑删除
    db.Delete(&product,1)




}