package main

import "github.com/gin-gonic/gin"

func main() {
	//default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "少林")
		address := c.DefaultQuery("address", "背景")
		//username := c.Query("username")
		//address := c.Query("address")
		c.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8081")
}
