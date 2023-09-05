package userbusiness

import (
	"context"
	usermodel "todo-app/module/user/model"
)

type DeleteUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
	) (*usermodel.User, error)

	DeleteUser(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteUser(ctx context.Context, condition map[string]interface{}) error {
	// step 1: Find User by conditions
	_, err := biz.store.FindUser(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to delete User
	if err := biz.store.DeleteUser(ctx, condition); err != nil {
		return err
	}

	return nil
}
