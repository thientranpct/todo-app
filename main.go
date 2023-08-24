package main

import (
	"net/http"

	todomodel "todo-app/module/item/model"
	todotrpt "todo-app/module/item/transport"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=Aa1234 dbname=todo_app port=5432 sslmode=disable TimeZone=Asia/Saigon"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&todomodel.TodoItem{})
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Todo app")
	})
	e.POST("items", todotrpt.HandleCreateItem(db))
	e.GET("items", todotrpt.HandleListItem(db))
	e.PUT("items/:id", todotrpt.HandleUpdateItem(db))
	e.GET("items/:id", todotrpt.HandleFindItem(db))
	e.DELETE("items/:id", todotrpt.HandleDeleteItem(db))
	e.Logger.Fatal(e.Start(":1323"))
}
