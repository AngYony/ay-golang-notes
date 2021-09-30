package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//实例化一个gin的server对象
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8021")

}
