package usertrpt

import (
	"net/http"
	"strconv"

	userbiz "todo-app/module/user/business"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// DeleteUser godoc
//
//	@Summary		Delete an user
//	@Description	Delete by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		200
//	@Router			/api/users/{id} [delete]
//	@Security		JWT
func HandleDeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewDeleteUserBiz(storage)

		if err := biz.DeleteUser(c.Request().Context(), map[string]interface{}{"id": id}); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, id)
	}
}
