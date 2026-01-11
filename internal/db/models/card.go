package models

import "time"

type Card struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	DueDate     time.Time `gorm:"not null" json:"dueDate"`
	Status      Status    `gorm:"size:50;not null" json:"status"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
