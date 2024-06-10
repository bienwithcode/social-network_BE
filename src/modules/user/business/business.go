package business

import (
	"context"
	"social-network/domain"
	"social-network/modules/auth/model"
	"social-network/utils"
)

type UserRepository interface {
	GetAuth(ctx context.Context, email, password string) (*domain.User, error)
	GetAuthUser(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context, authUserId string, paging *utils.Pagination, filter *model.Filter) ([]*domain.User, error)
	GetOnlineUsers(ctx context.Context, authUserId string) ([]*domain.User, error)
	GetNewMembers(ctx context.Context, authUserId string, paging *utils.Pagination) ([]*domain.User, error)
}

type business struct {
	userRepo UserRepository
}

func NewBusiness(userRepo UserRepository) *business {
	return &business{
		userRepo: userRepo,
	}
}

func (biz *business) GetAuth(ctx context.Context, email, password string) (*domain.User, error) {

	user, err := biz.userRepo.GetAuth(ctx, email, password)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (biz *business) GetAuthUser(ctx context.Context, id string) (*domain.User, error) {

	user, err := biz.userRepo.GetAuthUser(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (biz *business) GetUsers(ctx context.Context, authUserId string, paging *utils.Pagination, filter *model.Filter) ([]*domain.User, error) {

	user, err := biz.userRepo.GetUsers(ctx, authUserId, paging, filter)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (biz *business) GetOnlineUsers(ctx context.Context, authUserId string) ([]*domain.User, error) {

	user, err := biz.userRepo.GetOnlineUsers(ctx, authUserId)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (biz *business) GetNewMembers(ctx context.Context, authUserId string, paging *utils.Pagination) ([]*domain.User, error) {

	user, err := biz.userRepo.GetNewMembers(ctx, authUserId, paging)

	if err != nil {
		return nil, err
	}
	return user, nil
}
