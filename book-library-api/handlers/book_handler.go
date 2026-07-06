package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"book-library-api/database"
	"book-library-api/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if book.Title == "" || book.Author == "" || book.Genre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title, author and genre are required",
		})
		return
	}
	err := database.DB.QueryRow(
		`INSERT INTO books (title, author, published_year, genre)
	 VALUES ($1, $2, $3, $4)
	 RETURNING id`,
		book.Title,
		book.Author,
		book.PublishedYear,
		book.Genre,
	).Scan(&book.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		limitNum = 5
	}

	if limitNum > 100 {
		limitNum = 100
	}
	offset := (pageNum - 1) * limitNum

	rows, err := database.DB.Query(`
		SELECT id, title, author, published_year, genre
        FROM books
        ORDER BY id
        LIMIT $1 OFFSET $2
	`, limitNum, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	books := []models.Book{}

	for rows.Next() {
		var book models.Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.PublishedYear,
			&book.Genre,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	var book models.Book
	err = database.DB.QueryRow(
		`SELECT id, title, author, published_year, genre FROM books WHERE id=$1`,
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.PublishedYear,
		&book.Genre,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}
func UpdateBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if book.Title == "" || book.Author == "" || book.Genre == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title, author and genre are required",
		})
		return
	}
	result, err := database.DB.Exec(
		`UPDATE books
		 SET title = $1,
		     author = $2,
		     published_year = $3,
		     genre = $4
		 WHERE id = $5`,
		book.Title,
		book.Author,
		book.PublishedYear,
		book.Genre,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
	})
}

func DeleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid book ID",
		})
		return
	}
	result, err := database.DB.Exec(`DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}

func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title query parameter is required",
		})
		return
	}

	rows, err := database.DB.Query(
		`SELECT id, title, author, published_year, genre FROM books WHERE title ILIKE $1ORDER BY id`,
		"%"+title+"%",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {

		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.PublishedYear,
			&book.Genre,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}
