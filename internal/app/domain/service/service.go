package service

import "github.com/xabsvoid/eis/internal/app/domain/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
