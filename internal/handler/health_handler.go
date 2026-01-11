package handler

import (
	// "fmt"
	"net/http"
    "github.com/labstack/echo/v4"
)

func HealthCheckHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
    // fmt.Println("Healthcheck was calls...")
}