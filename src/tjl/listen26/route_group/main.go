package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "登陆成功",
	})
}
func submit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "点击了男",
	})
}
func read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "读取",
	})
}
func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}
	router.Run(":8088")
}
