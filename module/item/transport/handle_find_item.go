package todotrpt

import (
	"net/http"
	"strconv"

	todobiz "todo-app/module/item/business"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleFindItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		// setup dependencies
		storage := todostorage.NewPostgresStorage(db)
		biz := todobiz.NewFindTodoItemBiz(storage)

		data, err := biz.FindItem(c.Request().Context(), map[string]interface{}{"id": id})

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
