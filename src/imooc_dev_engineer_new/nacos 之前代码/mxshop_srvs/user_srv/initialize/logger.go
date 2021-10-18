package initialize

import "go.uber.org/zap"

func InitLogger() {
	// 设置替换全局Logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
