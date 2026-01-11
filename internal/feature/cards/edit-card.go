package cards

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/peeb01/todo-app/internal/db/models"
)

func UpdateCard(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var payload map[string]interface{}
		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "invalid request body",
			})
		}

		delete(payload, "id")
		delete(payload, "created_at")
		delete(payload, "updated_at")

		if err := db.Model(&models.Card{}).Where("id = ?", id).Updates(payload).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "failed to update card",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "card update successfully",
		})
	}
}

func UpdateCardAllValue(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var card models.Card

		if err := db.First(&card, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "Card not found",
			})
		}
		if err := c.Bind(&card); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "invalid request body",
			})
		}

		if err := db.Save(&card).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "failed to update card",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "card update successfully",
		})
	}
}
