package flow

import (
	"context"
	"testing"

	"github.com/neatflowcv/realtor/internal/pkg/repositories/core"
	"github.com/stretchr/testify/require"
)

func TestService_ListRealties(t *testing.T) {
	repository := core.NewRepository()
	service := NewService(repository)
	ctx := context.Background()
	realties, err := service.ListRealties(ctx, &ListRealtiesOptions{
		MaxDeposit: 1000,
		MaxRent:    70,
	})

	require.NoError(t, err)
	require.NotEmpty(t, realties)
}
