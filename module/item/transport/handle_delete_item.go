package todotrpt

import (
	"net/http"
	"strconv"

	todobiz "todo-app/module/item/business"
	todostorage "todo-app/module/item/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// DeleteItem godoc
//
//	@Summary		Delete a item
//	@Description	Delete by ID
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Item ID"
//	@Success		200
//	@Router			/api/items/{id} [delete]
//	@Security		JWT
func HandleDeleteItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		// setup dependencies
		storage := todostorage.NewPostgresStorage(db)
		biz := todobiz.NewDeleteTodoItemBiz(storage)

		if err := biz.DeleteItem(c.Request().Context(), map[string]interface{}{"id": id}); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, id)
	}
}
