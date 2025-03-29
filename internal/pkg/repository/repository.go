package repository

import (
	"github.com/neatflowcv/realtor/internal/pkg/domain"
)

type Options struct {
	MaxDeposit uint64 // 0 means no limit, 만원 단위
}

type Option func(*Options)

func WithMaxDeposit(maxDeposit uint64) Option {
	return func(o *Options) {
		o.MaxDeposit = maxDeposit
	}
}

type Repository interface {
	ListRealties(opts ...Option) ([]*domain.Realty, error)
}
