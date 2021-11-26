package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(context *gin.Context) {
		s := time.Now()

		context.Next()

		//使用zap记录日志
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			//获取间隔时间
			zap.Duration("elapsed", time.Now().Sub(s)))

	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())
		context.Next()

	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}

		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}

		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
