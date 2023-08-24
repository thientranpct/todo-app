package todostorage

import (
	"context"
	todomodel "todo-app/module/item/model"

	"gorm.io/gorm"
)

func (s *postgresStorage) FindItem(ctx context.Context, condition map[string]interface{}) (*todomodel.TodoItem, error) {
	var item todomodel.TodoItem

	if err := s.db.Where(condition).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, todomodel.ErrItemNotFound
		}
		return nil, err
	}
	return &item, nil
}
