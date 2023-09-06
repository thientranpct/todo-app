package usertrpt

import (
	"net/http"

	userbiz "todo-app/module/user/business"
	usermodel "todo-app/module/user/model"
	userstorage "todo-app/module/user/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ShowAccount godoc
//
//	@Summary		login
//	@Description	Login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			LoginPayload		body		usermodel.LoginPayload true	"LoginPayload"
//	@Success		200			{object}	usermodel.JwtCustomClaims
//	@Router			/login [post]
func HandleLogin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataItem usermodel.LoginPayload

		if err := c.Bind(&dataItem); err != nil {
			return err
		}
		// setup dependencies
		storage := userstorage.NewPostgresStorage(db)
		biz := userbiz.NewLoginBiz(storage)

		token, err := biz.Login(c.Request().Context(), map[string]interface{}{"email": dataItem.Email}, &dataItem)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Email or password is not correct")
		}

		return c.JSON(http.StatusOK, token)
	}
}
