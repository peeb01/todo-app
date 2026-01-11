package migration

import (
	"github.com/peeb01/todo-app/internal/db/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB){
	db.AutoMigrate(&models.Card{})
}