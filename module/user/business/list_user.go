package userbusiness

import (
	"context"
	usermodel "todo-app/module/user/model"
)

type ListUserStorage interface {
	ListUser(ctx context.Context) ([]usermodel.User, error)
}

type listBiz struct {
	store ListUserStorage
}

func NewListUserBiz(store ListUserStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListUser(ctx context.Context) ([]usermodel.User, error) {
	result, err := biz.store.ListUser(ctx)

	if err != nil {
		return nil, err
	}

	return result, err
}
