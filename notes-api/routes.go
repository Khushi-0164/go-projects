package main

import "github.com/gin-gonic/gin"

func setUpRoutes(router *gin.Engine) {
	router.GET("/", welcomeHandler)
	router.POST("/notes", createNoteHandler)
	router.GET("/notes", getAllNotesHandler)
	router.GET("/notes/:id", getNoteByIDHandler)
}
