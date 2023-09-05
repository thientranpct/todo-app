package userbusiness

import (
	"context"
	usermodel "todo-app/module/user/model"

	"golang.org/x/crypto/bcrypt"
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
	if err := data.Validate(); err != nil {
		return err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return err
	}
	data.Password = string(bytes)
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
