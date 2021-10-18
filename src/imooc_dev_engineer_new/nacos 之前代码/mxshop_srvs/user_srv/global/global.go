package global

import (
	"mxshop_srvs/user_srv/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

//
// func init() {
//
// 	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// 	dsn := "root:root123456@@tcp(59.110.216.174:7306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
//
// 	// 设置全局Logger，打印SQL
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
// 		logger.Config{
// 			SlowThreshold:             time.Second, // 慢 SQL 阈值
// 			LogLevel:                  logger.Info, // logger.Silent, // 日志级别
// 			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
// 			Colorful:                  true,        // 是否彩色打印
// 		},
// 	)
//
// 	// 全局模式
// 	var err error
// 	// 注意：此处不能通过:=赋值，否则将会新定义局部的DB变量，而非为全局DB赋值
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 		// 配置表名生成规则
// 		NamingStrategy: schema.NamingStrategy{
// 			// 设置以struct名为表名，而不是复数形式
// 			SingularTable: true,
// 		},
// 	})
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// 完成迁移，根据struct创建对应的表
// 	// _ = db.AutoMigrate(&model.User{})
// }
