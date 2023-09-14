package todotrpt

import (
	"net/http"
	"strconv"

	todobiz "todo-app/module/item/business"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ShowAccount godoc
//
//	@Summary		Show a todo item
//	@Description	get by ID
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Item ID"
//	@Success		200	{object}	todomodel.TodoItem
//	@Router			/api/items/{id} [get]
//	@Security		JWT
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
