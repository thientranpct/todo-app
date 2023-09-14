package usertrpt

import (
	"net/http"
	"strconv"
	"strings"

	userbiz "todo-app/module/user/business"
	usermodel "todo-app/module/user/model"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// UpdateUser godoc
//
//	@Summary		Update an user
//	@Description	Update an user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"User ID"
//	@Param			Payload	body		usermodel.User	true	"Payload"
//	@Success		200		{object}	usermodel.User
//	@Router			/api/users/{id} [put]
//	@Security		JWT
func HandleUpdateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return err
		}

		var dataItem usermodel.User

		if err := c.Bind(&dataItem); err != nil {
			return err
		}

		// preprocess title - trim all spaces
		dataItem.Email = strings.TrimSpace(dataItem.Email)

		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewUpdateUserBiz(storage)

		if err := biz.UpdateUser(c.Request().Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, dataItem)
	}
}
