package rpc

import (
	"context"
	"social-network/domain"
	"social-network/proto/pb"
)

type grpcServiceClient struct {
	usc pb.UserServiceClient
}

func NewGrpcServiceClient(usc pb.UserServiceClient) *grpcServiceClient {
	return &grpcServiceClient{usc: usc}
}

func (gsc *grpcServiceClient) GetAuth(ctx context.Context, email, password string) (*domain.User, error) {
	return nil, nil
}
