package main

import (
	"log"

	"blog-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Blog API is running!",
		})
	})

	router.Run(":8080")
}
