package flow

import (
	"testing"

	"github.com/neatflowcv/realtor/internal/pkg/repositories/core"
	"github.com/stretchr/testify/require"
)

func TestService_ListRealties(t *testing.T) {
	repository := core.NewRepository()
	service := NewService(repository)

	realties, err := service.ListRealties(&ListRealtiesOptions{
		MaxDeposit: 1000,
		MaxRent:    70,
	})

	require.NoError(t, err)
	require.NotEmpty(t, realties)
}
