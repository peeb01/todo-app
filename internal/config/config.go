package config

import (
    "os"
)

type Config struct {
    Port       string
    DBuser     string
    DBpassword string
    DBname     string
    DBport     string
    DBtimeZone string
}

func Load() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    dbuser := os.Getenv("DB_USER")
    dbpassword := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    dbport := os.Getenv("DB_PORT")
    timezone := os.Getenv("DB_TIMEZONE")

    return Config{
        Port:       port,
        DBuser:     dbuser,
        DBpassword: dbpassword,
        DBname:     dbname,
        DBport:     dbport,
        DBtimeZone: timezone,
    }
}
