package usertrpt

import (
	"net/http"
	"strconv"

	userbiz "todo-app/module/user/business"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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
