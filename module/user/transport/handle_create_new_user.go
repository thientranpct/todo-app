package usertrpt

import (
	"net/http"
	"strings"
	userbiz "todo-app/module/user/business"
	usermodel "todo-app/module/user/model"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleCreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataItem usermodel.User

		if err := c.Bind(&dataItem); err != nil {
			return err
		}

		// preprocess title - trim all spaces
		dataItem.Email = strings.TrimSpace(dataItem.Email)

		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewCreateUserBiz(storage)

		if err := biz.CreateNewUser(c.Request().Context(), &dataItem); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, dataItem)
	}
}
