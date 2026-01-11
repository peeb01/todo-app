package cards

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/peeb01/todo-app/internal/db/models"
)

func GetAllCards(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cards []models.Card
		if err := db.Find(&cards).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to fetch card",
			})
		}
		return c.JSON(http.StatusOK, cards)
	}
}

func GetCardByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var card models.Card
		if err := db.First(&card, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Card not found",
			})
		}
		
		return c.JSON(http.StatusOK, card)
	}
}