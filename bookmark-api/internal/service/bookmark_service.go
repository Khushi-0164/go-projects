package service

import (
	"bookmark-api/internal/models"
)

type BookmarkRepository interface {
	Create(bookmark *models.Bookmark) error
	FindAll(page, limit int, tag string) ([]models.Bookmark, int64, error)
	Delete(id uint) error
}

type BookmarkService struct {
	repo BookmarkRepository
}

func NewBookmarkService(repo BookmarkRepository) *BookmarkService {
	return &BookmarkService{repo: repo}
}

func (s *BookmarkService) Create(title, url, tags string) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{Title: title, URL: url, Tags: tags}
	if err := s.repo.Create(bookmark); err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarkService) List(page, limit int, tag string) ([]models.Bookmark, int64, error) {
	return s.repo.FindAll(page, limit, tag)
}

func (s *BookmarkService) Delete(id uint) error {
	return s.repo.Delete(id)
}
