package todomodel

import (
	"errors"
	"time"
)

var (
	ErrTitleCannotBeBlank        = errors.New("title can not be blank")
	ErrItemNotFound              = errors.New("item not found")
	ErrCannotUpdateCompletedItem = errors.New("can not update completed item")
)

type TodoItem struct {
	Id        int        `json:"id" gorm:"column:id" swaggerignore:"true"`
	Title     string     `json:"title" gorm:"column:title"`
	Status    string     `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;" swaggerignore:"true"`
}

func (TodoItem) TableName() string { return "todo_items" }

func (item TodoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}
	return nil
}
