package main

import (
	"log"

	"book-library-api/database"

	"book-library-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")

}
