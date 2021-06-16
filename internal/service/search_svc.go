package service

import (
	"context"
	"gprc-template/pkg/pb/generated/api"
)

type searchImpl struct {}

func NewSearchImpl() api.SearchServer {
	return &searchImpl{}
}

func (s searchImpl) ByQuery(ctx context.Context, request *api.SearchRequest) (*api.SearchResponse, error) {
	return &api.SearchResponse{
		Id: "1",
	}, nil
}


