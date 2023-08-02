package todostorage

import (
	"context"
	todomodel "todo-app/module/item/model"
)

func (s *postgresStorage) CreateItem(ctx context.Context, data *todomodel.TodoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
