package handlers

import (
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkHandler struct {
	DB *gorm.DB
}

func NewLinkHandler(db *gorm.DB) *LinkHandler {
	return &LinkHandler{DB: db}
}

type createLinkRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
	var req createLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("user_id")
	var code string
	for attempts := 0; attempts < 5; attempts++ {
		generated, err := utils.GenerateShortCode(7)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var existing models.Link
		h.DB.Where("short_code = ?", generated).First(&existing)
		if err := h.DB.Where("short_code = ?", generated).First(&existing).Error; err == gorm.ErrRecordNotFound {
			code = generated
			break
		}
	}
	if code == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate unique short code"})
		return
	}
	link := models.Link{
		ShortCode:   code,
		OriginalURL: req.URL,
		UserID:      userID,
	}
	if err := h.DB.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create link"})
		return
	}
	c.JSON(http.StatusCreated, link)
}

func (h *LinkHandler) ListLinks(c *gin.Context) {
	userID := c.GetUint("user_id")
	var links []models.Link
	if err := h.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&links).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch links"})
		return
	}
	c.JSON(http.StatusOK, links)
}

func (h *LinkHandler) DeleteLink(c *gin.Context) {
	id := c.Param("id")
	userId := c.GetUint("user_id")
	role, _ := c.Get("role")

	var link models.Link
	if err := h.DB.First(&link, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}
	if link.UserID != userId && role != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not allowed to delete this link"})
		return
	}
	if err := h.DB.Delete(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete link"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "link deleted successfully"})
}

func (h *LinkHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	var link models.Link
	if err := h.DB.Where("short_code = ?", code).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "short link not found"})
		return
	}

	h.DB.Model(&link).UpdateColumn("clicks", gorm.Expr("clicks + 1"))

	c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
}
