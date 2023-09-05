package userbusiness

import (
	"context"
	usermodel "todo-app/module/user/model"

	"golang.org/x/crypto/bcrypt"
)

type UpdateUserStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error)
	UpdateUser(ctx context.Context, condition map[string]interface{}, dataUpdate *usermodel.User) error
}

type updateBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateUser(ctx context.Context, condition map[string]interface{}, dataUpdate *usermodel.User) error {
	if dataUpdate.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(dataUpdate.Password), 14)
		if err != nil {
			return err
		}
		dataUpdate.Password = string(bytes)
	}
	if err := biz.store.UpdateUser(ctx, condition, dataUpdate); err != nil {
		return err
	}
	return nil
}
