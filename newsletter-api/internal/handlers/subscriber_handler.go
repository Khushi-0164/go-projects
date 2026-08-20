package handlers

import (
	"errors"
	"net/http"
	"newsletter-api/internal/service"

	"github.com/gin-gonic/gin"
)

type SubscriberHandler struct {
	service *service.SubscriberService
}

func NewSubscriberHandler(s *service.SubscriberService) *SubscriberHandler {
	return &SubscriberHandler{service: s}
}

type subscribeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (h *SubscriberHandler) Subscribe(c *gin.Context) {
	var req subscribeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriber, err := h.service.Subscribe(req.Email)
	if err != nil {
		if errors.Is(err, service.ErrAlreadySubscribed) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to subscribe"})
		return
	}
	c.JSON(http.StatusCreated, subscriber)
}
