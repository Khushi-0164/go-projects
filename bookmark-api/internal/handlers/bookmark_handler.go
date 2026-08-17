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
	bookmarks, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookmarks"})
		return
	}
	c.JSON(http.StatusOK, bookmarks)
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
