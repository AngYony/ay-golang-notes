package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")

		// 告诉浏览器，允许使用哪些方法
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PATCH,PUT")
		// 允许哪些Header
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		method := c.Request.Method

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

	}
}
