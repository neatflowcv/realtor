package core

import (
	"context"
	"fmt"
	"log"
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

func (r *Repository) ListRealties(ctx context.Context, opts ...repository.Option) ([]*domain.Realty, error) {
	var options repository.Options
	for _, opt := range opts {
		opt(&options)
	}

	codes := r.client.ListCodes()

	var lists []*zigbang.List
	for i, code := range codes {
		log.Printf("search code %v %v (%v/%v)", code.ID, code.Location, i+1, len(codes))
		list, err := r.client.GetCatalogList(context.Background(), code.ID, options.MaxDeposit, options.MaxRent)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list.List...)
	}

	var realties []*domain.Realty
	for _, item := range lists {
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
