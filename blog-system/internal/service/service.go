package service

import "project/internal/repository"

type Service struct {
	Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Repository: repository}
}
