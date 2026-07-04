package main

import (
	"database/sql"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

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

	err = db.QueryRow(
		"INSERT INTO notes(title, content) VALUES($1, $2) RETURNING id",
		note.Title,
		note.Content,
	).Scan(&note.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create note",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
		"note":    note,
	})
}

func getAllNotesHandler(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, content FROM notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve notes",
		})
		return
	}
	defer rows.Close()

	var allNotes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve notes",
			})
			return
		}
		allNotes = append(allNotes, note)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve notes",
		})
		return
	}
	c.JSON(http.StatusOK, allNotes)
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

	var note Note

	err = db.QueryRow(
		"SELECT id, title, content FROM notes WHERE id = $1",
		id,
	).Scan(&note.ID, &note.Title, &note.Content)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve note",
		})
		return
	}

	c.JSON(http.StatusOK, note)
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
	var updatedNote Note
	err = c.ShouldBindJSON(&updatedNote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	result, err := db.Exec(
		"UPDATE notes SET title = $1, content = $2 WHERE id = $3",
		updatedNote.Title,
		updatedNote.Content,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update note",
		})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update note",
		})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	updatedNote.ID = id
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

	result, err := db.Exec(
		"DELETE FROM notes WHERE id=$1",
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete note",
		})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete note",
		})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}
