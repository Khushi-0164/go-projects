package handlers

import (
	"habit-tracker/models"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateHabit(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	habit := models.Habit{
		UserId:    userID,
		Name:      input.Name,
		CreatedAt: time.Now(),
	}

	if err := models.DB.Create(&habit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create habit"})
		return
	}
	c.JSON(http.StatusCreated, habit)
}
func CheckIn(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	habitID := c.Param("id")

	var habit models.Habit
	if err := models.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "habit not found"})
		return
	}

	today := time.Now().Truncate(24 * time.Hour)

	var existing models.CheckIn
	err := models.DB.Where("habit_id = ? AND date = ?", habit.ID, today).First(&existing).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "already checked in today"})
		return
	}

	checkIn := models.CheckIn{HabitID: habit.ID, Date: today}
	models.DB.Create(&checkIn)

	c.JSON(http.StatusCreated, gin.H{"message": "checked in", "date": today})
}
func ListHabits(c *gin.Context) {
	userId := c.MustGet("user_id").(uint)
	var habits []models.Habit
	if err := models.DB.Where("user_id=?", userId).Find(&habits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch habits"})
		return
	}
	type HabitResponse struct {
		ID             uint    `json:"id"`
		Name           string  `json:"name"`
		CurrentStreak  int     `json:"current_streak"`
		CompletionRate float64 `json:"completion_rate"`
	}
	var response []HabitResponse
	for _, h := range habits {
		var checkIns []models.CheckIn
		models.DB.Where("habit_id = ?", h.ID).Order("date desc").Find(&checkIns)

		streak := calculateStreak(checkIns)
		completion := calculateCompletionRate(h.CreatedAt, len(checkIns))

		response = append(response, HabitResponse{
			ID:             h.ID,
			Name:           h.Name,
			CurrentStreak:  streak,
			CompletionRate: completion,
		})
	}
	c.JSON(http.StatusOK, response)
}
func calculateStreak(checkIns []models.CheckIn) int {
	if len(checkIns) == 0 {
		return 0
	}

	today := time.Now().Truncate(24 * time.Hour)
	yesterday := today.Add(-24 * time.Hour)

	// streak only counts if the most recent check-in was today or yesterday
	mostRecent := checkIns[0].Date
	if !mostRecent.Equal(today) && !mostRecent.Equal(yesterday) {
		return 0
	}

	streak := 1
	expectedDate := mostRecent.Add(-24 * time.Hour)

	for i := 1; i < len(checkIns); i++ {
		if checkIns[i].Date.Equal(expectedDate) {
			streak++
			expectedDate = expectedDate.Add(-24 * time.Hour)
		} else {
			break
		}
	}

	return streak
}
func calculateCompletionRate(createdAt time.Time, totalCheckIns int) float64 {
	daysSinceCreated := int(time.Since(createdAt).Hours()/24) + 1
	if daysSinceCreated <= 0 {
		return 0
	}
	rate := (float64(totalCheckIns) / float64(daysSinceCreated)) * 100
	return math.Round(rate*100) / 100
}
func GetHabit(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	habitID := c.Param("id")

	var habit models.Habit
	if err := models.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "habit not found"})
		return
	}

	var checkIns []models.CheckIn
	models.DB.Where("habit_id = ?", habit.ID).Order("date desc").Find(&checkIns)

	c.JSON(http.StatusOK, gin.H{
		"habit":     habit,
		"check_ins": checkIns,
	})
}
func DeleteHabit(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	habitID := c.Param("id")

	var habit models.Habit
	if err := models.DB.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "habit not found"})
		return
	}

	models.DB.Where("habit_id = ?", habit.ID).Delete(&models.CheckIn{})
	models.DB.Delete(&habit)

	c.JSON(http.StatusOK, gin.H{"message": "habit deleted"})
}
