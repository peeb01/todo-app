package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var DB *gorm.DB

/*
ConnectDatabase establishes a connection to a PostgreSQL database.

Parameters:
  host      : Database host address
  user      : Database username
  password  : Database user password
  dbname    : Database name
  port      : Database port number
  sslmode   : SSL mode setting (disable, require, verify-full, etc.)
  timezone  : Database server timezone (e.g. Asia/Bangkok)

  This function should be called once at application startup 
*/
func ConnectDatabase(host, user, password, dbname, port, sslmode, timezone string) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected")
}