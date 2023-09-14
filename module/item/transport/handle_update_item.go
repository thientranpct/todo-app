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

// UpdateItem godoc
//
//	@Summary		Update item
//	@Description	Update item
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Item ID"
//	@Param			Payload	body		todomodel.TodoItem	true	"Payload"
//	@Success		200		{object}	todomodel.TodoItem
//	@Router			/api/items/{id} [put]
//	@Security		JWT
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
