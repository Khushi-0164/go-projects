package handlers

import (
	"booking-api/internal/models"
	"booking-api/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourceHandler struct {
	repo *repository.ResourceRepository
}

func NewResourceHandler(repo *repository.ResourceRepository) *ResourceHandler {
	return &ResourceHandler{repo: repo}
}

type createResourceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var req createResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resource := &models.Resource{Name: req.Name, Description: req.Description}
	if err := h.repo.Create(resource); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create resource"})
		return
	}
	c.JSON(http.StatusCreated, resource)
}

func (h *ResourceHandler) ListResources(c *gin.Context) {
	resources, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch resources"})
		return
	}
	c.JSON(http.StatusOK, resources)
}
