package services

import "gitlab.com/voxe-analytics/internal/core/repository"

type Service struct {
}

func New(repos repository.Store) *Service {
	return &Service{}
}
