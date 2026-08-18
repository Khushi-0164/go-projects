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

func (r *BookmarkRepository) FindAll(page, limit int, tag string) ([]models.Bookmark, int64, error) {
	var bookmarks []models.Bookmark
	var total int64

	query := r.DB.Model(&models.Bookmark{})
	if tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&bookmarks).Error

	return bookmarks, total, err
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
