package server

import (
	"context"
	"math/rand"

	"github.com/rashad-j/go-grpc-search-svc/config"
)

type SearchService interface {
	InsertRandomData(max int32) int32
	search(context.Context, int32) int32
	insert(context.Context, int32) int32
	delete(context.Context, int32) int32
}

type searchService struct {
	data []int32
}

func NewSearchService(cfg *config.Config) SearchService {
	return &searchService{
		data: make([]int32, 0, cfg.MaxSize),
	}
}

// InsertRandomData inserts random data into the data slice
func (s *searchService) InsertRandomData(max int32) int32 {
	for i := int32(0); i < max; i++ {
		// generate random number
		n := rand.Int31n(max)
		s.insert(context.Background(), n)
	}
	return int32(len(s.data))
}

// search performs a binary search on the data slice and returns the position of the target number
// or -1 if the target number is not found
func (s *searchService) search(_ context.Context, target int32) int32 {
	var left int32 = 0
	var right int32 = int32(len(s.data)) - 1

	for left <= right {
		mid := (left + right) / 2

		if s.data[mid] == target {
			return mid
		}

		if s.data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// insert inserts the target number into the data slice and returns the position of the inserted number (ordered asc)
func (s *searchService) insert(_ context.Context, target int32) int32 {
	var left int32 = 0
	var right int32 = int32(len(s.data)) - 1

	for left <= right {
		mid := (left + right) / 2

		if s.data[mid] == target {
			return mid
		}

		if s.data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	s.data = append(s.data, 0)
	copy(s.data[left+1:], s.data[left:])
	s.data[left] = target

	return left
}

// delete deletes the target number from the data slice and returns the position of the deleted number
func (s *searchService) delete(_ context.Context, target int32) int32 {
	var left int32 = 0
	var right int32 = int32(len(s.data)) - 1

	for left <= right {
		mid := (left + right) / 2

		if s.data[mid] == target {
			s.data = append(s.data[:mid], s.data[mid+1:]...)
			return mid
		}

		if s.data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
