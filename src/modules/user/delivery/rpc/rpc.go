package rpc

import (
	"context"
	"social-network/domain"
	"social-network/proto/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Business interface {
	GetAuth(ctx context.Context, email, password string) (*domain.User, error)
	GetAuthUser(ctx context.Context, id string) (*domain.User, error)
}

type grpcService struct {
	pb.UnimplementedUserServiceServer
	business Business
}

func NewGrpcService(business Business) *grpcService {
	return &grpcService{business: business}
}

func (gs *grpcService) GetAuth(ctx context.Context, req *pb.GetAuthRequest) (*pb.GetAuthResponse, error) {
	user, err := gs.business.GetAuth(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	pbUser := &pb.User{
		Id:            user.Id,
		Email:         user.Email,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
	}
	return &pb.GetAuthResponse{
		User: pbUser,
	}, nil
}

func (gs *grpcService) GetAuthUser(ctx context.Context, req *pb.GetAuthUserRequest) (*pb.GetAuthUserResponse, error) {
	user, err := gs.business.GetAuthUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAuthUserResponse{
		User: &pb.User{
			Id:            user.Id,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
			Banned:        user.Banned,
			FacebookId:    user.FacebookId,
			GoogleId:      user.GoogleId,
			GithubId:      user.GithubId,
			IsOnline:      user.IsOnline,
			Posts:         user.Posts,
			Likes:         user.Likes,
			Comments:      user.Comments,
			Followers:     user.Followers,
			Following:     user.Following,
			Messages:      user.Messages,
			Notifications: user.Notifications,
			FullName:      user.FullName,
			Email:         user.Email,
			CreatedAt:     timestamppb.New(*user.CreatedAt),
			UpdatedAt:     timestamppb.New(*user.UpdatedAt),
		},
	}, nil
}
