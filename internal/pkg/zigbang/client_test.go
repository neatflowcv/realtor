package zigbang_test

import (
	"context"
	"testing"

	"github.com/neatflowcv/realtor/internal/pkg/zigbang"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCatalogList(t *testing.T) {
	t.Parallel()
	client := zigbang.NewClient()
	ctx := context.Background()

	catalog, err := client.GetCatalogList(ctx, "11680106", 0)

	require.NoError(t, err)
	require.NotNil(t, catalog)
	require.NotEmpty(t, catalog.List)
	require.Len(t, catalog.List, 10)
}

func TestClient_ListCodes(t *testing.T) {
	t.Parallel()
	client := zigbang.NewClient()

	codes := client.ListCodes()

	require.NotEmpty(t, codes)
}
