package userstorage

import (
	"context"
	usermodel "cmc/module/users/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *usermodel.User) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}