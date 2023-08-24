package todostorage

import (
	"context"
	todomodel "todo-app/module/item/model"
)

func (s *postgresStorage) ListItem(ctx context.Context) ([]todomodel.TodoItem, error) {
	var result []todomodel.TodoItem

	if err := s.db.Table(todomodel.TodoItem{}.TableName()).
		Find(&result).Error; err != nil {

		return nil, err
	}
	return result, nil
}
