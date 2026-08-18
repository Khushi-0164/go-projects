package handlers

import (
	"bookmark-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookmarkHandler struct {
	service *service.BookmarkService
}

func NewBookingHandler(s *service.BookmarkService) *BookmarkHandler {
	return &BookmarkHandler{service: s}
}

type createBookmarkRequest struct {
	Title string `json:"title" binding:"required" `
	URL   string `json:"url" binding:"required,url" `
	Tags  string `json:"tags" `
}

func (h *BookmarkHandler) Create(c *gin.Context) {
	var req createBookmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookmark, err := h.service.Create(req.Title, req.URL, req.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bookmark"})
		return
	}
	c.JSON(http.StatusCreated, bookmark)
}

func (h *BookmarkHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 5
	}
	tag := c.Query("tag")

	bookmarks, total, err := h.service.List(page, limit, tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookmarks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        bookmarks,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *BookmarkHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete bookmark"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bookmark deleted"})
}
