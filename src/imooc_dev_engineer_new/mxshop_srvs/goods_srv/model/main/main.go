package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"mxshop_srvs/goods_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string {
	myMd5 := md5.New()
	_, _ = io.WriteString(myMd5, code)
	return hex.EncodeToString(myMd5.Sum(nil))

}

func main() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root123456@@tcp(59.110.216.174:7306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局Logger，打印SQL
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 是否彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		// 配置表名生成规则
		NamingStrategy: schema.NamingStrategy{
			// 设置以struct名为表名，而不是复数形式
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	// 完成迁移，根据struct创建对应的表
	_ = db.AutoMigrate(
		&model.Category{},
		&model.Brands{},
		&model.GoodsCategoryBrand{},
		&model.Banner{},
		&model.Goods{},
	)

}
