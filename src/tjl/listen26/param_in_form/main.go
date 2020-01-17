package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(200, gin.H{
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
