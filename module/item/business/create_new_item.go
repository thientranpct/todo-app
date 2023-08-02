package todobusiness

import (
	"context"
	todomodel "todo-app/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.TodoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateTodoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.TodoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.Status = "Doing"

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
