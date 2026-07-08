package handlers

import (
	"blog-api/database"
	"blog-api/models"
	"blog-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = string(hashedPassword)

	err = database.DB.QueryRow(
		`INSERT INTO users (username, email, password)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully",
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func LoginUser(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User

	err := database.DB.QueryRow(
		`SELECT id, username, email, password
     FROM users
     WHERE email = $1`,
		loginRequest.Email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginRequest.Password),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}
