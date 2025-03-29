package repository

import (
	"context"

	"github.com/neatflowcv/realtor/internal/pkg/domain"
)

type Options struct {
	MaxDeposit uint64 // 보증금, 0 means no limit, 만원 단위
	MaxRent    uint64 // 월세, 0 means no limit, 만원 단위
}

type Option func(*Options)

func WithMaxDeposit(maxDeposit uint64) Option {
	return func(o *Options) {
		o.MaxDeposit = maxDeposit
	}
}

func WithMaxRent(maxRent uint64) Option {
	return func(o *Options) {
		o.MaxRent = maxRent
	}
}

type Repository interface {
	ListRealties(ctx context.Context, opts ...Option) ([]*domain.Realty, error)
}
