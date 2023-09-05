package userstorage

import (
	"context"
	usermodel "todo-app/module/user/model"

	"gorm.io/gorm"
)

func (s *postgresStorage) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error) {
	var user usermodel.User

	if err := s.db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, usermodel.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
