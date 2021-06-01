package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	// server.POST("/test", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	page := c.DefaultQuery("page", "0")
	// 	name := c.PostForm("name")
	// 	message := c.PostForm("message")

	// 	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

	// })

	server.Run(":8080")
}
