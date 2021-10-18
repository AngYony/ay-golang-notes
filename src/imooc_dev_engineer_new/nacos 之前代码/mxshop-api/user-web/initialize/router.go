package initialize

import (
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// 配置允许跨域访问
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router

}
