package main

import (
	"os"

	"github.com/rashad-j/go-grpc-search-svc/config"
	"github.com/rashad-j/go-grpc-search-svc/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// zerolog basic config
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	// create searcher
	searcher := server.NewSearchService(cfg)
	searcher = server.NewLoggerService(searcher)

	// generate some random data
	searcher.InsertRandomData(100_000)

	// create server
	server.MakeSearchServer(searcher)
}
