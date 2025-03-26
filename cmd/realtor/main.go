package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/neatflowcv/realtor/internal/app/flow"
	"github.com/neatflowcv/realtor/internal/pkg/repositories/core"
)

func main() {
	repo := core.NewRepository()
	service := flow.NewService(repo)
	realties, err := service.ListRealties()
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
			"source: %s(%s), transaction: %s\n",
			realty.SourceKind(),
			realty.SourceID(),
			realty.TransactionKind(),
		)
		if err != nil {
			log.Printf("failed to write realty: %v", err)
			break
		}
	}
}
