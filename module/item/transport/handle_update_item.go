package todotrpt

import (
	"net/http"
	"strconv"
	"strings"

	todobiz "todo-app/module/item/business"
	todomodel "todo-app/module/item/model"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleUpdateItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		var dataItem todomodel.TodoItem

		if err := c.Bind(&dataItem); err != nil {
			return err
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		// setup dependencies
		storage := todostorage.NewPostgresStorage(db)
		biz := todobiz.NewUpdateTodoItemBiz(storage)

		if err := biz.UpdateItem(c.Request().Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, dataItem)
	}
}
