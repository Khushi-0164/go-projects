package repository

import (
	"bookmark-api/internal/models"

	"gorm.io/gorm"
)

type BookmarkRepository struct {
	DB *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{DB: db}
}

func (r *BookmarkRepository) Create(bookmark *models.Bookmark) error {
	return r.DB.Create(bookmark).Error
}

func (r *BookmarkRepository) FindAll() ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark
	err := r.DB.Order("created_at desc").Find(&bookmarks).Error
	return bookmarks, err
}

func (r *BookmarkRepository) FindBtId(id uint) (*models.Bookmark, error) {
	var bookmark models.Bookmark
	if err := r.DB.First(&bookmark, id).Error; err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (r *BookmarkRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Bookmark{}, id).Error
}
