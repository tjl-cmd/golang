package main

import "github.com/gin-gonic/gin"

func testHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", testHandle)
	r.Run(":9090")
}
