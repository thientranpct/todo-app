package userstorage

import (
	"context"
	usermodel "todo-app/module/user/model"
)

func (s *postgresStorage) UpdateUser(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *usermodel.User,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
