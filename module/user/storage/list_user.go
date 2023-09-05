package userstorage

import (
	"context"
	usermodel "todo-app/module/user/model"
)

func (s *postgresStorage) ListUser(ctx context.Context) ([]usermodel.User, error) {
	var result []usermodel.User

	if err := s.db.Table(usermodel.User{}.TableName()).
		Find(&result).Error; err != nil {

		return nil, err
	}
	return result, nil
}
