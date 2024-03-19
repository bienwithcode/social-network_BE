package business

import (
	"context"
	"social-network/domain"
	"social-network/modules/auth/model"
	"social-network/utils"
	"time"
)

type AuthRepository interface {
}

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

func (biz *business) Login(ctx context.Context, data *model.AuthEmailPassword) (*model.TokenResponse, error) {

	authData, err := biz.userRepo.GetAuth(ctx, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	claimsData := make(map[string]interface{})
	claimsData["auth"] = &authData

	expiredAt := time.Hour * 24

	tokenStr, expSecs, err := utils.IssueToken(claimsData, expiredAt)

	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{
		AccessToken: model.Token{
			Token:     tokenStr,
			ExpiredIn: expSecs,
		},
	}, nil
}
