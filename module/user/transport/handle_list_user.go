package usertrpt

import (
	"net/http"

	userbiz "todo-app/module/user/business"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleListUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewListUserBiz(storage)

		data, err := biz.ListUser(c.Request().Context())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
