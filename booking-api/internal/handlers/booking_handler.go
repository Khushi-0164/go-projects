package handlers

import (
	"booking-api/internal/service"
	"booking-api/internal/utils"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service *service.BookingService
}

func NewBookingHandler(s *service.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

type createBookingRequest struct {
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	userID := c.GetUint("user_id")
	resourceID := utils.ParseUint(c.Param("id"))

	var req createBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking, err := h.service.CreateBooking(resourceID, userID, req.StartTime, req.EndTime)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrSlotUnavailable):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrInvalidTimeRange):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrResourceNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, booking)
}
func (h *BookingHandler) ListForResource(c *gin.Context) {
	resourceID := utils.ParseUint(c.Param("id"))

	bookings, err := h.service.ListForResource(resourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
func (h *BookingHandler) ListMyBookings(c *gin.Context) {
	userID := c.GetUint("user_id")
	bookings, err := h.service.ListForUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
