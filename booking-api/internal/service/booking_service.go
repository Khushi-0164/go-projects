package service

import (
	"booking-api/internal/models"
	"booking-api/internal/repository"
	"errors"
	"time"
)

var (
	ErrSlotUnavailable  = errors.New("time slot is unavailable")
	ErrInvalidTimeRange = errors.New("start_time must be before end_time")
	ErrResourceNotFound = errors.New("resource not found")
)

type BookingService struct {
	bookingRepo  *repository.BookingRepository
	resourceRepo *repository.ResourceRepository
}

func NewBookingService(bookingRepo *repository.BookingRepository, resourceRepo *repository.ResourceRepository) *BookingService {
	return &BookingService{bookingRepo: bookingRepo, resourceRepo: resourceRepo}
}

func (s *BookingService) CreateBooking(resourceID, userID uint, start, end time.Time) (*models.Booking, error) {
	if !start.Before(end) {
		return nil, ErrInvalidTimeRange
	}

	if _, err := s.resourceRepo.FindByID(resourceID); err != nil {
		return nil, ErrResourceNotFound
	}

	booking := &models.Booking{
		ResourceID: resourceID,
		UserID:     userID,
		StartTime:  start,
		EndTime:    end,
	}

	if err := s.bookingRepo.CreateIfAvailavle(booking); err != nil {
		if errors.Is(err, repository.ErrSlotUnavailable) {
			return nil, ErrSlotUnavailable
		}
		return nil, err
	}
	return booking, nil
}

func (s *BookingService) ListForResource(resourceID uint) ([]models.Booking, error) {
	return s.bookingRepo.FindByResource(resourceID)
}

func (s *BookingService) ListForUser(userID uint) ([]models.Booking, error) {
	return s.bookingRepo.FindByUser(userID)
}
