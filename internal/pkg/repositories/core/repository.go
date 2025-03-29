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
		zigbangCode := code.ID[:len(code.ID)-2]
		log.Printf("search code %v %v (%v/%v)", zigbangCode, code.Location, i+1, len(codes))
		page := &zigbang.Pagination{
			Offset: 0,
			Limit:  200,
		}
		out, err := r.client.GetCatalogList(ctx, zigbangCode, options.MaxDeposit, options.MaxRent, page)
		if err != nil {
			return nil, err
		}
		log.Println("count", out.Count)
		lists = append(lists, out.List...)
		for page.Limit == len(out.List) {
			page.Offset += page.Limit
			out, err = r.client.GetCatalogList(ctx, code.ID, options.MaxDeposit, options.MaxRent, page)
			if err != nil {
				return nil, err
			}
			lists = append(lists, out.List...)
		}
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
