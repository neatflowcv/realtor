package repository

import "github.com/neatflowcv/realtor/internal/pkg/domain"

type Repository interface {
	ListRealties() ([]*domain.Realty, error)
}
