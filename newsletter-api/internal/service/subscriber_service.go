package service

import (
	"errors"
	"newsletter-api/internal/models"
	"newsletter-api/internal/repository"
	"newsletter-api/internal/worker"
)

var ErrAlreadySubscribed = errors.New("email already subscribed")

type SubscriberService struct {
	repo *repository.SubscriberRepository
	pool *worker.Pool
}

func NewSubscriberService(repo *repository.SubscriberRepository, pool *worker.Pool) *SubscriberService {
	return &SubscriberService{repo: repo, pool: pool}
}

func (s *SubscriberService) Subscribe(email string) (*models.Subscriber, error) {
	if s.repo.ExistsByEmail(email) {
		return nil, ErrAlreadySubscribed
	}
	subscriber := &models.Subscriber{Email: email}
	if err := s.repo.Create(subscriber); err != nil {
		return nil, err
	}
	s.pool.Enqueue(worker.Job{Email: email})
	return subscriber, nil
}
