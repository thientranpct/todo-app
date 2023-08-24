package todobusiness

import (
	"context"
	todomodel "todo-app/module/item/model"
)

type FindTodoItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*todomodel.TodoItem, error)
}

type findBiz struct {
	store FindTodoItemStorage
}

func NewFindTodoItemBiz(store FindTodoItemStorage) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) FindItem(ctx context.Context, condition map[string]interface{}) (*todomodel.TodoItem, error) {
	item, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return nil, err
	}

	return item, nil
}
