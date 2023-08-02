package todomodel

import (
	"errors"
	"time"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type TodoItem struct {
	Id        int        `json:"id" gorm:"column:id"`
	Title     string     `json:"title" gorm:"column:title"`
	Status    string     `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (item TodoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}
	return nil
}
