package account

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Printf("total cast time:%d us", latency/1000)
	}
}
