package service

import (
	"context"

	"invoice-saas/internal/repository"
)

type SummaryService struct {
	repo *repository.CachedSummaryRepository
}

func NewSummaryService(repo *repository.CachedSummaryRepository) *SummaryService {
	return &SummaryService{repo: repo}
}

func (s *SummaryService) GetSummary(ctx context.Context, orgID uint) (*repository.OrgSummary, error) {
	return s.repo.GetSummary(ctx, orgID)
}
