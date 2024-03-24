package business

import (
	"context"
	"social-network/domain"
)

type UserRepository interface {
	GetAuth(ctx context.Context, email, password string) (*domain.User, error)
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
