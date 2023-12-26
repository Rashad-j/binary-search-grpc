package mock_search

// unit test cases for MockSearchServiceServer interface

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rashad-j/go-grpc-search-svc/rpc/search"
)

func TestMockSearchServiceServer_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearchServiceServer := NewMockSearchServiceServer(ctrl)

	ctx := context.Background()
	req := &search.DeleteRequest{
		Number: 1,
	}

	mockSearchServiceServer.EXPECT().Delete(ctx, req).Return(&search.DeleteResponse{Position: 0}, nil)

	deleteResponse, err := mockSearchServiceServer.Delete(ctx, req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if deleteResponse.Position != 0 {
		t.Errorf("Unexpected position: %v", deleteResponse.Position)
	}
}

func TestMockSearchServiceServer_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearchServiceServer := NewMockSearchServiceServer(ctrl)

	ctx := context.Background()
	req := &search.InsertRequest{
		Number: 1,
	}

	mockSearchServiceServer.EXPECT().Insert(ctx, req).Return(&search.InsertResponse{Position: 0}, nil)

	insertResponse, err := mockSearchServiceServer.Insert(ctx, req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if insertResponse.Position != 0 {
		t.Errorf("Unexpected position: %v", insertResponse.Position)
	}
}

func TestMockSearchServiceServer_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearchServiceServer := NewMockSearchServiceServer(ctrl)

	ctx := context.Background()
	req := &search.SearchRequest{
		Number: 1,
	}
	mockSearchServiceServer.EXPECT().Search(ctx, req).Return(&search.SearchResponse{Position: 0}, nil)
	searchResponse, err := mockSearchServiceServer.Search(ctx, req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if searchResponse.Position != 0 {
		t.Errorf("Unexpected position: %v", searchResponse.Position)
	}

	req = &search.SearchRequest{
		Number: 2,
	}
	mockSearchServiceServer.EXPECT().Search(ctx, req).Return(&search.SearchResponse{Position: 1}, nil)
	searchResponse, err = mockSearchServiceServer.Search(ctx, req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if searchResponse.Position != 1 {
		t.Errorf("Unexpected position: %v", searchResponse.Position)
	}
}

func TestMockSearchServiceServer_mustEmbedUnimplementedSearchServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSearchServiceServer := NewMockSearchServiceServer(ctrl)

	mockSearchServiceServer.EXPECT().mustEmbedUnimplementedSearchServiceServer()

	mockSearchServiceServer.mustEmbedUnimplementedSearchServiceServer()
}
