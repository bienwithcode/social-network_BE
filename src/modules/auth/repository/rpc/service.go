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
		Role:          res.User.Role,
		EmailVerified: res.User.EmailVerified,
	}, nil
}

func (gsc *grpcServiceClient) GetAuthUser(ctx context.Context, id string) (*domain.User, error) {
	res, err := gsc.usc.GetAuthUser(ctx, &pb.GetAuthUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	createAt := res.User.CreatedAt.AsTime()
	updatedAt := res.User.UpdatedAt.AsTime()
	return &domain.User{
		Id:            res.User.Id,
		Role:          res.User.Role,
		EmailVerified: res.User.EmailVerified,
		Banned:        res.User.Banned,
		FacebookId:    res.User.FacebookId,
		GoogleId:      res.User.GoogleId,
		GithubId:      res.User.GithubId,
		IsOnline:      res.User.IsOnline,
		Posts:         res.User.Posts,
		Likes:         res.User.Likes,
		Comments:      res.User.Comments,
		Followers:     res.User.Followers,
		Following:     res.User.Following,
		Messages:      res.User.Messages,
		Notifications: res.User.Notifications,
		FullName:      res.User.FullName,
		Email:         res.User.Email,
		CreatedAt:     &createAt,
		UpdatedAt:     &updatedAt,
	}, nil
}
