package router

import (
    "net/http"
    "github.com/peeb01/todo-app/internal/handler"
)

func New() http.Handler {
    mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HealthCheckHandler)
    mux.HandleFunc("/healthchecks", handler.HealthCheckHandler)
    return mux
}
