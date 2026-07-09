package handlers

import (
	"blog-api/database"
	"blog-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	blog.UserID = c.GetInt("user_id")

	err := database.DB.QueryRow(
		`INSERT INTO blogs (title, content, user_id)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at`,
		blog.Title,
		blog.Content,
		blog.UserID,
	).Scan(
		&blog.ID,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create blog",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    blog,
	})
}

func GetAllBlogs(c *gin.Context) {
	rows, err := database.DB.Query(
		`SELECT id, title, content, user_id, created_at, updated_at
	 FROM blogs`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch blogs",
		})
		return
	}
	defer rows.Close()

	var blogs []models.Blog

	for rows.Next() {
		var blog models.Blog

		err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.UserID,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read blog",
			})
			return
		}

		blogs = append(blogs, blog)
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}
func GetBlogByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid blog ID",
		})
		return
	}
	var blog models.Blog
	err = database.DB.QueryRow(
		`SELECT id, title, content, user_id, created_at, updated_at
        FROM blogs
        WHERE id = $1	`,
		id,
	).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.UserID,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"blog": blog,
	})
}
func UpdateBlog(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid blog ID",
		})
		return
	}
	userID := c.GetInt("user_id")

	var ownerID int
	err = database.DB.QueryRow(
		`SELECT user_id
	 FROM blogs
	 WHERE id = $1`,
		id,
	).Scan(&ownerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}
	if ownerID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not authorized to update this blog",
		})
		return
	}

	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = database.DB.Exec(
		`UPDATE blogs
	 SET title = $1,
	     content = $2,
	     updated_at = CURRENT_TIMESTAMP
	 WHERE id = $3`,
		blog.Title,
		blog.Content,
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update blog",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
	})
}
func DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid blog ID",
		})
		return
	}

	userID := c.GetInt("user_id")

	var ownerID int
	err = database.DB.QueryRow(
		`SELECT user_id
		 FROM blogs
		 WHERE id = $1`,
		id,
	).Scan(&ownerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Blog not found",
		})
		return
	}

	if ownerID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You are not authorized to delete this blog",
		})
		return
	}

	_, err = database.DB.Exec(
		`DELETE FROM blogs
		 WHERE id = $1`,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete blog",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}
