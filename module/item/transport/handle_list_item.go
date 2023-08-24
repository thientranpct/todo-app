package todotrpt

import (
	"net/http"

	todobiz "todo-app/module/item/business"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleListItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// setup dependencies
		storage := todostorage.NewPostgresStorage(db)
		biz := todobiz.NewListTodoItemBiz(storage)

		data, err := biz.ListItem(c.Request().Context())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
