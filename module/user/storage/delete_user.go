package userstorage

import (
	"context"
	usermodel "todo-app/module/user/model"
)

func (s *postgresStorage) DeleteUser(ctx context.Context, condition map[string]interface{}) error {
	if err := s.db.Table(usermodel.User{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
