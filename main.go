package main

import (
	"net/http"

	todomodel "todo-app/module/item/model"
	todotrpt "todo-app/module/item/transport"
	usermodel "todo-app/module/user/model"
	usertrpt "todo-app/module/user/transport"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "todo-app/docs"
)

//	@title			Swagger API Todo app
//	@version		1.0
//	@description	Todo app.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host						localhost:1323
//
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
func main() {
	dsn := "host=localhost user=postgres password=Aa1234 dbname=todo_app port=5432 sslmode=disable TimeZone=Asia/Saigon"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&todomodel.TodoItem{})
	db.AutoMigrate(&usermodel.User{})
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Todo app")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	restricted := e.Group("api")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(usermodel.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	restricted.Use(echojwt.WithConfig(config))
	restricted.POST("items", todotrpt.HandleCreateItem(db))
	restricted.GET("items", todotrpt.HandleListItem(db))
	restricted.PUT("items/:id", todotrpt.HandleUpdateItem(db))
	restricted.GET("items/:id", todotrpt.HandleFindItem(db))
	restricted.DELETE("items/:id", todotrpt.HandleDeleteItem(db))

	restricted.POST("users", usertrpt.HandleCreateUser(db))
	restricted.GET("users", usertrpt.HandleListUser(db))
	restricted.PUT("users/:id", usertrpt.HandleUpdateUser(db))
	restricted.GET("users/:id", usertrpt.HandleFindUser(db))
	restricted.DELETE("users/:id", usertrpt.HandleDeleteUser(db))

	e.POST("login", usertrpt.HandleLogin(db))
	e.Logger.Fatal(e.Start(":1323"))
}
