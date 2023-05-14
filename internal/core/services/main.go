package services

import "gitlab.com/greatsoft/xif-backend/internal/core/repository"

type Service struct {
}

func New(repos repository.Store) *Service {
	return &Service{}
}
