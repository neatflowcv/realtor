package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/neatflowcv/realtor/internal/pkg/domain"
	"github.com/neatflowcv/realtor/internal/pkg/repository"
	"github.com/neatflowcv/realtor/internal/pkg/zigbang"
)

var _ repository.Repository = (*Repository)(nil)

type Repository struct {
	client *zigbang.Client
}

func NewRepository() *Repository {
	return &Repository{
		client: zigbang.NewClient(),
	}
}

func (r *Repository) ListRealties() ([]*domain.Realty, error) {
	codes := r.client.ListCodes()

	// TODO: code를 순회해야 한다.
	list, err := r.client.GetCatalogList(context.Background(), codes[0].ID)
	if err != nil {
		return nil, err
	}

	var realties []*domain.Realty
	for _, item := range list.List {
		id := strconv.FormatUint(uint64(item.AreaHoID), 10) //nolint:gosec
		source := domain.NewRealtySource(domain.SourceKindZigbang, id)
		area := domain.NewArea(item.SizeContractM2, item.SizeM2)
		var transactionKind domain.TransactionKind
		switch item.TranType {
		case "trade":
			transactionKind = domain.TransactionKindTrade
		case "charter":
			transactionKind = domain.TransactionKindJeonse
		case "rental":
			transactionKind = domain.TransactionKindRent
		default:
			panic(fmt.Sprintf("unknown transaction type: %s", item.TranType))
		}
		builder := domain.NewRealtyBuilder(source, transactionKind).
			Deposit(item.DepositMin).
			Rent(item.RentMin).
			Area(area)

		realties = append(realties, builder.Build())
	}

	return realties, nil
}
