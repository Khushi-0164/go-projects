package service

import (
	"bookmark-api/internal/models"
	"bookmark-api/internal/repository"
)

type BookmarkService struct {
	repo *repository.BookmarkRepository
}

func NewBookingService(repo *repository.BookmarkRepository) *BookmarkService {
	return &BookmarkService{repo: repo}
}

func (s *BookmarkService) Create(title, url, tags string) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{Title: title, URL: url, Tags: tags}
	if err := s.repo.Create(bookmark); err != nil {
		return nil, err
	}
	return bookmark, nil
}

func (s *BookmarkService) List() ([]models.Bookmark, error) {
	return s.repo.FindAll()
}

func (s *BookmarkService) Delete(id uint) error {
	return s.repo.Delete(id)
}
