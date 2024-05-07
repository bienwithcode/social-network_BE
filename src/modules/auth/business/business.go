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
	GetAuthorize(ctx context.Context, email, password string) (*domain.User, error)
	GetAuthUser(ctx context.Context, id string) (*domain.User, error)
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

	authData, err := biz.userRepo.GetAuthorize(ctx, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	claimsData := make(map[string]interface{})
	claimsData["auth"] = &authData

	expiredAt := time.Hour * 24

	tokenStr, _, err := utils.IssueToken(claimsData, expiredAt)

	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{
		Token: tokenStr,
		User: model.User{
			Id:    authData.Id,
			Email: authData.Email,
			Role:  authData.Role,
		},
	}, nil
}

func (biz *business) GetAuthUser(ctx context.Context, data *model.AuthUserId) (*domain.User, error) {

	authData, err := biz.userRepo.GetAuthUser(ctx, data.Id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Id:            authData.Id,
		Role:          authData.Role,
		EmailVerified: authData.EmailVerified,
		Banned:        authData.Banned,
		FacebookId:    authData.FacebookId,
		GoogleId:      authData.GoogleId,
		GithubId:      authData.GithubId,
		IsOnline:      authData.IsOnline,
		Posts:         authData.Posts,
		Likes:         authData.Likes,
		Comments:      authData.Comments,
		Followers:     authData.Followers,
		Following:     authData.Following,
		Messages:      authData.Messages,
		Notifications: authData.Notifications,
		FullName:      authData.FullName,
		Email:         authData.Email,
		CreatedAt:     authData.CreatedAt,
		UpdatedAt:     authData.UpdatedAt,
	}, nil
}
