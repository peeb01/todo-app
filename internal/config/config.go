package config

import (
	"os"
	"strings"
)

type Config struct {
	Port        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_TIMEZONE string
	DB_SSLMODE  string
	JWT_SECRET  string
}

/*
Load reads application configuration from environment variables.

Environment Variables:

	PORT         : Application port (default: 5000 if empty)
	DB_HOST      : Database host address
	DB_USER      : Database username
	DB_PASSWORD  : Database user password
	DB_NAME      : Database name
	DB_PORT      : Database port
	DB_TIMEZONE  : Database timezone (e.g. Asia/Bangkok)
	DB_SSLMODE   : Database SSL mode (disable, require, etc.)
	JWT_SECRET   : Secret key for signing JSON Web Tokens

	Returns:
	Config struct containing all application and database configuration values.
*/
func Load() Config {
	port := strings.TrimSpace(os.Getenv("APP_PORT"))
	if port == "" {
		port = "8080"
	}

	dbhost := strings.TrimSpace(os.Getenv("DB_HOST"))
	dbuser := strings.TrimSpace(os.Getenv("DB_USER"))
	dbpassword := strings.TrimSpace(os.Getenv("DB_PASSWORD"))
	dbname := strings.TrimSpace(os.Getenv("DB_NAME"))
	dbport := strings.TrimSpace(os.Getenv("DB_PORT"))
	timezone := strings.TrimSpace(os.Getenv("DB_TIMEZONE"))
	sslmode := strings.TrimSpace(os.Getenv("DB_SSLMODE"))
	jwtSecret := strings.TrimSpace(os.Getenv("JWT_SECRET"))

	return Config{
		Port:        port,
		DB_HOST:     dbhost,
		DB_USER:     dbuser,
		DB_PASSWORD: dbpassword,
		DB_NAME:     dbname,
		DB_PORT:     dbport,
		DB_TIMEZONE: timezone,
		DB_SSLMODE:  sslmode,
		JWT_SECRET:  jwtSecret,
	}
}
