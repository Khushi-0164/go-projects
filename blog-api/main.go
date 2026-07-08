package main

import (
	"log"

	"blog-api/database"
	"blog-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
