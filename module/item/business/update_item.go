package todobusiness

import (
	"context"
	todomodel "todo-app/module/item/model"
)

type UpdateTodoItemStorage interface {
	FindItem(ctx context.Context, condition map[string]interface{}) (*todomodel.TodoItem, error)
	UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *todomodel.TodoItem) error
}

type updateBiz struct {
	store UpdateTodoItemStorage
}

func NewUpdateTodoItemBiz(store UpdateTodoItemStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *todomodel.TodoItem) error {
	item, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	if item.Status == "Completed" {
		return todomodel.ErrCannotUpdateCompletedItem
	}

	if err := biz.store.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}
	return nil
}
