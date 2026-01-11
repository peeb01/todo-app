package cards

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/peeb01/todo-app/internal/db/models"
)

func NewCard(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		card := new(models.Card)
		if err := c.Bind(card); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid request body",
			})
		}
		if err := db.Create(card).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "failed to create card",
			})
		}
		return c.JSON(http.StatusCreated, card)
	}
}
