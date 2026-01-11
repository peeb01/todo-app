package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/peeb01/todo-app/internal/handler"
	"github.com/peeb01/todo-app/internal/feature/cards"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	api.GET("", handler.HealthCheckHandler)
	api.GET("/", handler.HealthCheckHandler)

	api.GET("/cards", cards.GetAllCards(db))
	api.GET("/cards/:id", cards.GetCardByID(db))
	api.POST("/cards", cards.NewCard(db))
	api.PUT("/cards/:id", cards.UpdateCard(db))
	api.PUT("/cards/:id/all", cards.UpdateCardAllValue(db))
	api.DELETE("/cards/:id", cards.DeleteCard(db))

	return e
}
