package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	ut "github.com/go-playground/universal-translator"

	myvalidator "mxshop-api/user-web/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

func main() {

	// 初始化Logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()

	// 初始化routers
	Router := initialize.Routers()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的mymobile要和forms/user.go中的mobile的tag对应
		_ = v.RegisterValidation("mymobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mymobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mymobile", "{0} 非法的手机号码！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mymobile", fe.Field())
			return t
		})
	}

	port := global.ServerConfig.Port

	// zap.S()是对Sugar的进一步封装
	zap.S().Infof("启动服务器，端口：%d", port)

	err := Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
