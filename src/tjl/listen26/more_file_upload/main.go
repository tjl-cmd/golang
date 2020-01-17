package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files upload", len(files)),
		})
	})
	router.Run(":9090")
}
