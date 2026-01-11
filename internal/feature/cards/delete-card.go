package cards

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/peeb01/todo-app/internal/db/models"
)

func DeleteCard(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var card models.Card
		if err := db.First(&card, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Card not found",
			})
		}
		if err := db.Delete(&card).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Failed to delete card",
			
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Card deleted successfully",
		})
	}
}
