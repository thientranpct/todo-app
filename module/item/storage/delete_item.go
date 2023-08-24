package todostorage

import (
	"context"
	todomodel "todo-app/module/item/model"
)

func (s *postgresStorage) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	if err := s.db.Table(todomodel.TodoItem{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
