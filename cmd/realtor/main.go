package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/neatflowcv/realtor/internal/app/flow"
	"github.com/neatflowcv/realtor/internal/pkg/repositories/core"
)

func main() {
	repo := core.NewRepository()
	service := flow.NewService(repo)
	ctx := context.Background()
	realties, err := service.ListRealties(ctx, &flow.ListRealtiesOptions{
		MaxDeposit: 1000,
		MaxRent:    70,
	})
	if err != nil {
		log.Fatalf("failed to list realties: %v", err)
	}
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		err := writer.Flush()
		if err != nil {
			log.Printf("failed to flush writer: %v", err)
		}
	}()
	for _, realty := range realties {
		_, err := fmt.Fprintf(
			writer,
			"source: %s(%s), transaction: %s, deposit: %d, rent: %d, total area: %.2f, net area: %.2f\n",
			realty.SourceKind(),
			realty.SourceID(),
			realty.TransactionKind(),
			realty.Deposit(),
			realty.Rent(),
			realty.TotalArea(),
			realty.NetArea(),
		)
		if err != nil {
			log.Printf("failed to write realty: %v", err)
			break
		}
	}
}
