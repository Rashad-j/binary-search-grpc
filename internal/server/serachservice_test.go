package server

// test cases for searchService
import (
	"context"
	"testing"

	"github.com/rashad-j/go-grpc-search-svc/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSearchService_InsertRandomData(t *testing.T) {
	// arrange
	cfg := &config.Config{
		MaxSize: 100,
	}

	// act
	s := NewSearchService(cfg)
	l := s.InsertRandomData(100)

	// assert
	assert.Equal(t, int32(l), int32(len(s.(*searchService).data)))
}

func TestSearchService_search(t *testing.T) {
	// arrange
	cfg := &config.Config{
		MaxSize: 100,
	}

	// act
	s := NewSearchService(cfg)
	s.insert(context.Background(), 0)
	s.insert(context.Background(), 99)
	s.insert(context.Background(), 88)

	// assert
	assert.Equal(t, int32(0), s.search(context.Background(), 0))
	assert.Equal(t, int32(2), s.search(context.Background(), 99))
	assert.Equal(t, int32(1), s.search(context.Background(), 88))
	assert.Equal(t, int32(-1), s.search(context.Background(), 100))
}

func TestSearchService_insert(t *testing.T) {
	// arrange
	cfg := &config.Config{
		MaxSize: 100,
	}

	// act
	s := NewSearchService(cfg)
	s.insert(context.Background(), 0)
	s.insert(context.Background(), 99)

	// assert
	assert.Equal(t, int32(0), s.(*searchService).data[0])
	assert.Equal(t, int32(99), s.(*searchService).data[1])
}

func TestSearchService_delete(t *testing.T) {
	// arrange
	cfg := &config.Config{
		MaxSize: 100,
	}

	// act
	s := NewSearchService(cfg)
	s.insert(context.Background(), 0)
	s.insert(context.Background(), 99)
	s.delete(context.Background(), 0)
	s.delete(context.Background(), 99)

	// assert
	assert.Equal(t, 0, len(s.(*searchService).data))
}
