package client

import (
	"github.com/pkg/errors"
	"github.com/rashad-j/go-grpc-search-svc/config"
	"github.com/rashad-j/go-grpc-search-svc/rpc/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MakeSearchServiceClient creates a new gRPC client and returns the client, connection, and error
func MakeSearchServiceClient(cfg *config.Config) (search.SearchServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to dial server")
	}

	client := search.NewSearchServiceClient(conn)

	return client, conn, nil
}
