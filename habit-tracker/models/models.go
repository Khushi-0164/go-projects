package models

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"-"`
}

type Habit struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type CheckIn struct {
	ID      uint      `json:"id" gorm:"primaryKey"`
	HabitID uint      `json:"habit_id"`
	Date    time.Time `json:"date"`
}
