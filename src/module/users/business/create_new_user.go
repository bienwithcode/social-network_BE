package userbusiness

import (
	usermodel "cmc/module/users/model"
	"context"
	"errors"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.User) error
}

type createBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewUser(ctx context.Context, data *usermodel.User) error {

	if data.Username == "" {
		return errors.New("title can not be blank")
	}
	// do not allow "finished" status when creating a new task
	data.Status = 1 // set to default

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}
	return errors.New("")
}
