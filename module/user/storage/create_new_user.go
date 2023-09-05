package userstorage

import (
	"context"
	usermodel "todo-app/module/user/model"
)

func (s *postgresStorage) CreateUser(ctx context.Context, data *usermodel.User) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
