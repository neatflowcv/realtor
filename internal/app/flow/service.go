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
}

type ListRealtiesOption func(*ListRealtiesOptions)

func WithMaxDeposit(maxDeposit uint64) ListRealtiesOption {
	return func(o *ListRealtiesOptions) {
		o.MaxDeposit = maxDeposit
	}
}

func (s *Service) ListRealties(opts ...ListRealtiesOption) ([]*domain.Realty, error) {
	var options ListRealtiesOptions
	for _, opt := range opts {
		opt(&options)
	}
	return s.repository.ListRealties(repository.WithMaxDeposit(options.MaxDeposit))
}
