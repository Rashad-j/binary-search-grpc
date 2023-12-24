package server

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/rashad-j/go-grpc-search-svc/config"
	"github.com/rashad-j/go-grpc-search-svc/rpc/search"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type searchServer struct {
	searcher SearchService
	search.UnimplementedSearchServiceServer
}

func NewSearchServer(searcher SearchService) *searchServer {
	return &searchServer{
		searcher: searcher,
	}
}

func (s *searchServer) Search(ctx context.Context, req *search.SearchRequest) (*search.SearchResponse, error) {
	target := req.GetTargetNumber()

	position := s.searcher.search(ctx, target)

	return &search.SearchResponse{
		Position: position,
	}, nil
}

func (s *searchServer) Insert(ctx context.Context, req *search.InsertRequest) (*search.InsertResponse, error) {
	target := req.GetNumber()

	position := s.searcher.insert(ctx, target)

	return &search.InsertResponse{
		Position: position,
	}, nil
}

func (s *searchServer) Delete(ctx context.Context, req *search.DeleteRequest) (*search.DeleteResponse, error) {
	target := req.GetNumber()

	position := s.searcher.delete(ctx, target)

	return &search.DeleteResponse{
		Position: position,
	}, nil
}

func MakeSearchServer(searcher SearchService) error {
	// read config
	cfg, err := config.ReadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	// create server
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	search.RegisterSearchServiceServer(s, NewSearchServer(searcher))

	// Enable reflection
	reflection.Register(s)

	log.Info().Msgf("gRPC listening on %s", cfg.Addr)

	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

// loggingInterceptor is a unary interceptor that logs each incoming RPC.
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// set requestID in context
	requestID := xid.New().String()
	ctx = context.WithValue(ctx, "requestID", requestID)

	// metrics & logging
	log.Info().Str("requestID", requestID).Str("method", info.FullMethod).Msg("starting RPC request")
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msg("finished RPC request")
	}(time.Now())

	// calling the actual handler to process the RPC
	resp, err := handler(ctx, req)

	return resp, err
}
