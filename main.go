package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kevin021788/project/models"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var login models.User
		//把POST JSON 绑定结构体
		if err := c.ShouldBindJSON(&login); err != nil {
			//如果错误停止访问
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}

		c.JSON(http.StatusOK,gin.H{"status":"200","list":login})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}