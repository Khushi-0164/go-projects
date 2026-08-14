package repository

import (
	"booking-api/internal/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrSlotUnavailable = errors.New("time slot is unavailable")

type BookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (r *BookingRepository) CreateIfAvailavle(booking *models.Booking) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var existing []models.Booking
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("resource_id = ? AND start_time < ? AND end_time > ?",
				booking.ResourceID, booking.EndTime, booking.StartTime).
			Find(&existing).Error
		if err != nil {
			return err
		}
		if len(existing) > 0 {
			return ErrSlotUnavailable
		}

		return tx.Create(booking).Error
	})
}

func (r *BookingRepository) FindByResource(resourceID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.Where("resource_id=?", resourceID).Order("start_time asc").Find(&bookings).Error
	return bookings, err
}
func (r *BookingRepository) FindByUser(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.DB.Where("user_id = ?", userID).Order("start_time asc").Find(&bookings).Error
	return bookings, err
}
