package flow

import (
	"github.com/neatflowcv/realtor/internal/pkg/domain"
	"github.com/neatflowcv/realtor/internal/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListRealties() ([]*domain.Realty, error) {
	return s.repository.ListRealties()
}
