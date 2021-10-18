package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mxshop_srvs/user_srv/model"
	"os"
	"time"

	"github.com/anaskhan96/go-password-encoder"

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
	dsn := "root:root123456@@tcp(59.110.216.174:7306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	// // 完成迁移，根据struct创建对应的表
	// _ = db.AutoMigrate(&model.User{})

	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("AAA%d", i),
			Mobile:   fmt.Sprintf("123456789%d", i),
			PassWord: newPassword,
		}
		db.Save(&user)
	}

	// // Using custom options
	// options := &password.Options{10, 100, 32, sha512.New}
	// salt, encodedPwd := password.Encode("generic password", options)
	//
	// newPassword := fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
	// println(len(newPassword))
	// pwdinfo := strings.Split(newPassword, "$")
	// salt = pwdinfo[1]
	// encodedPwd = pwdinfo[2]
	//
	// check := password.Verify("generic password", salt, encodedPwd, options)
	// fmt.Println(check) // true

	// b := generateSalt(5)
	// fmt.Println(string(b))
	// fmt.Println(hex.EncodeToString(b))
}

//b5b4cabdd6f2e0a65196088b04e8642e
//b3383dd70b89e6feb99a49b9e4f6991c
//7183dd4a9e4979cfcee58bee8c73772a

func generateSalt(length int) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	salt := make([]byte, length)
	rand.Read(salt)
	for key, val := range salt {
		salt[key] = alphanum[val%byte(len(alphanum))]
	}
	return salt
}
