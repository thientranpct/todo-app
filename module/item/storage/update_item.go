package todostorage

import (
	"context"
	todomodel "todo-app/module/item/model"
)

func (s *postgresStorage) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *todomodel.TodoItem,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
