package server

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

// loggerService is a decorator for the SearchService interface that logs the target and result of each search
type loggerService struct {
	next SearchService
}

// add a compile time check to ensure that loggerService implements the SearchService interface
var _ SearchService = (*loggerService)(nil)

func NewLoggerService(next SearchService) *loggerService {
	return &loggerService{
		next: next,
	}
}

func (s *loggerService) InsertRandomData(max int32) int32 {
	log.Info().Msgf("inserting random data for max: %d", max)
	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Dur("duration", duration).Msgf("insert random data complete for max: %d", max)
	}(time.Now())

	return s.next.InsertRandomData(max)
}

func (s *loggerService) search(ctx context.Context, target int32) int32 {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}

	log.Info().Str("requestID", requestID).Msgf("searching target: %d", target)

	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msgf("search complete for target: %d", target)
	}(time.Now())

	return s.next.search(ctx, target)
}

func (s *loggerService) insert(ctx context.Context, target int32) int32 {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}

	log.Info().Str("requestID", requestID).Msgf("inserting target: %d", target)

	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msgf("insert complete for target: %d", target)
	}(time.Now())

	return s.next.insert(ctx, target)
}

func (s *loggerService) delete(ctx context.Context, target int32) int32 {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}

	log.Info().Str("requestID", requestID).Msgf("deleting target: %d", target)

	defer func(start time.Time) {
		duration := time.Since(start)
		log.Info().Str("requestID", requestID).Dur("duration", duration).Msgf("delete complete for target: %d", target)
	}(time.Now())

	return s.next.delete(ctx, target)
}
