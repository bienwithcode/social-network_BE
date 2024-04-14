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

func (gsc *grpcServiceClient) GetAuthorize(ctx context.Context, email, password string) (*domain.User, error) {
	res, err := gsc.usc.GetAuth(ctx, &pb.GetAuthRequest{Email: email, Password: password})
	if err != nil {
		return nil, err
	}
	return &domain.User{
		Id:            res.User.Id,
		Email:         res.User.Email,
		Username:      res.User.Username,
		Role:          res.User.Role,
		EmailVerified: res.User.EmailVerified,
		// CreatedAt: res.User.CreatedAt,
		// UpdatedAt: res.User.UpdatedAt,
	}, nil
}
