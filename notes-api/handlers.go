package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var notes []Note

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Notes API!",
	})
}
func createNoteHandler(c *gin.Context) {
	var note Note
	err := c.ShouldBindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	notes = append(notes, note)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
		"note":    note,
		"count":   len(notes),
	})
}

func getAllNotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, notes)
}

func getNoteByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid note ID",
		})
		return
	}
	if id < 0 || id >= len(notes) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	c.JSON(http.StatusOK, notes[id])

}

func updateNoteByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid note ID",
		})
		return
	}
	if id < 0 || id >= len(notes) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	var updatedNote Note
	err = c.ShouldBindJSON(&updatedNote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	notes[id] = updatedNote
	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated successfully",
		"note":    updatedNote,
	})
}
func deleteNoteByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid note ID",
		})
		return
	}
	if id < 0 || id >= len(notes) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	notes = append(notes[:id], notes[id+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}
