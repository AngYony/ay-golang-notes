# zap

用于记录日志。

两种模式：

- sugar 模式：使用方便，由于内部需要启用反射机制，因此性能不如Logger模式，但依然远强于其他日志组件，推荐使用。
- Logger 模式：性能最好，一般用于追求极致性能的场景。



```go
package main

import (
    "go.uber.org/zap"
)

func main()  {
    logger, _ := zap.NewProduction() // 生产环境
    // logger,_ := zap.NewDevelopment() // 开发环境
    defer logger.Sync() // flushes buffer, if any
    url := "https://imooc.com"

    // sugar 模式：使用方便，但性能不如sugar，因为要启用反射机制
    sugar := logger.Sugar()
    sugar.Infow("failed to fetch URL",
        // Structured context as loosely typed key-value pairs.
        "url", url, // 类似于map形式
        "attempt", 3,
    )
    sugar.Infof("Failed to fetch URL: %s", url)

    // logger模式：追求极致性能
    logger.Info("日志信息",
        zap.String("url", url),
        zap.Int("nums",3))

}
```



## zap 将日志输出到文件中

```go
package main

import (
	"time"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myproject.log",
		"stderr",
		"stdout",
	}
	return cfg.Build()
}

func main() {
	// logger, _ := zap.NewProduction()
	logger, err := NewLogger()
	if err != nil {
		panic(err)
		// panic("初始化logger失败")
	}
	su := logger.Sugar()
	defer su.Sync()
	url := "https://imooc.com"
	su.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

```

