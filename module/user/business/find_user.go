package userbusiness

import (
	"context"
	usermodel "todo-app/module/user/model"
)

type FindUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodel.User, error)
}

type findBiz struct {
	store FindUserStorage
}

func NewFindUserBiz(store FindUserStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error) {
	user, err := biz.store.FindUser(ctx, condition)

	if err != nil {
		return nil, err
	}

	return user, nil
}
