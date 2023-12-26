package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/rashad-j/go-grpc-search-svc/internal/client"
	"github.com/rashad-j/go-grpc-search-svc/internal/config"
	"github.com/rashad-j/go-grpc-search-svc/rpc/search"
)

func main() {
	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	searchClient, conn, err := client.MakeSearchServiceClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create client")
	}
	defer conn.Close()

	// insert
	insertResponse, err := searchClient.Insert(context.Background(), &search.InsertRequest{
		Number: 42,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to insert")
	}
	log.Info().Int32("position", insertResponse.Position).Msg("insert result")

	// search
	searchResponse, err := searchClient.Search(context.Background(), &search.SearchRequest{
		Number: 42,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to search")
	}
	log.Info().Int32("position", searchResponse.Position).Msg("search result")

	// delete
	deleteResponse, err := searchClient.Delete(context.Background(), &search.DeleteRequest{
		Number: 42,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to delete")
	}

	log.Info().Int32("position", deleteResponse.Position).Msg("delete result")

	// search again
	searchResponse, err = searchClient.Search(context.Background(), &search.SearchRequest{
		Number: 42,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to search")
	}
	log.Info().Int32("position", searchResponse.Position).Msg("search again result")
}
