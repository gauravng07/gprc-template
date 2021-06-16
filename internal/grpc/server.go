package grpc

import "google.golang.org/grpc"
import middleware "github.com/grpc-ecosystem/go-grpc-middleware"

func CreateServerWithLogV1(opts []grpc.ServerOption) *grpc.Server  {
	opts = append(opts, grpc.UnaryInterceptor(middleware.ChainUnaryServer()))
	server := grpc.NewServer(opts...)
	return server
}