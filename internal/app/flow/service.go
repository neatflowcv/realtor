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

type ListRealtiesOptions struct {
	MaxDeposit uint64
	MaxRent    uint64
}

func (s *Service) ListRealties(options *ListRealtiesOptions) ([]*domain.Realty, error) {
	return s.repository.ListRealties(repository.WithMaxDeposit(options.MaxDeposit), repository.WithMaxRent(options.MaxRent))
}
