package rpc

import (
	"context"
	"social-network/domain"
	"social-network/proto/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Business interface {
	GetAuth(ctx context.Context, email, password string) (*domain.User, error)
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
		Username:      user.Username,
		Email:         user.Email,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		CreatedAt:     timestamppb.New(*user.CreatedAt),
		UpdatedAt:     timestamppb.New(*user.UpdatedAt),
	}
	return &pb.GetAuthResponse{
		User: pbUser,
	}, nil
}
