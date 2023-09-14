package todotrpt

import (
	"net/http"
	"strings"
	todobiz "todo-app/module/item/business"
	todomodel "todo-app/module/item/model"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateItem godoc
//
//	@Summary		Add a todo item
//	@Description	Add a todo item
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			Payload	body		todomodel.TodoItem	true	"Payload"
//	@Success		200		{object}	todomodel.TodoItem
//	@Router			/api/items [post]
//	@Security		JWT
func HandleCreateItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataItem todomodel.TodoItem

		if err := c.Bind(&dataItem); err != nil {
			return err
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		// setup dependencies
		storage := todostorage.NewPostgresStorage(db)
		biz := todobiz.NewCreateTodoItemBiz(storage)

		if err := biz.CreateNewItem(c.Request().Context(), &dataItem); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, dataItem)
	}
}
