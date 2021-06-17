package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gprc-template/pkg/pb/generated/api"
)

type searchImpl struct {}

func NewSearchImpl() api.SearchServer {
	return &searchImpl{}
}

func (s searchImpl) ByQuery(ctx context.Context, request *api.SearchRequest) (*api.SearchResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Errorf("failed to get metadata")
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}

	reqId := md["x-request-id"]
	if len(reqId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing 'x-request-id' header")
	}
	fmt.Print(reqId)

	header := metadata.New(map[string]string{
		"x-request-id": "res-123",
	})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return nil, status.Errorf(codes.Internal, "unable to send header")
	}

	return &api.SearchResponse{
		Id: "1",
	}, nil
}


