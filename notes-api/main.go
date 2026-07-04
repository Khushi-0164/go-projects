package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	connectDB()
	router := gin.Default()

	setUpRoutes(router)

	router.Run(":8080")
}
