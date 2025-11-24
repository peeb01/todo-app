package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	UserID	 	uint `gorm:"primaryKey"`
	Name		string
	Email     	string `gorm:"unique"`

}

