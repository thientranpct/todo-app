package usertrpt

import (
	"net/http"
	"strconv"

	userbiz "todo-app/module/user/business"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ShowAccount godoc
//
//	@Summary		Show an user
//	@Description	get by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	usermodel.User
//	@Router			/api/users/{id} [get]
//	@Security		JWT
func HandleFindUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewFindUserBiz(storage)

		data, err := biz.FindUser(c.Request().Context(), map[string]interface{}{"id": id})

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
