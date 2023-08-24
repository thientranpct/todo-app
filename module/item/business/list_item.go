package todobusiness

import (
	"context"
	todomodel "todo-app/module/item/model"
)

type ListTodoItemStorage interface {
	ListItem(ctx context.Context) ([]todomodel.TodoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListTodoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItem(ctx context.Context) ([]todomodel.TodoItem, error) {
	result, err := biz.store.ListItem(ctx)

	if err != nil {
		return nil, err
	}

	return result, err
}
