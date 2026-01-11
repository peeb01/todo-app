# Todo App API

This is a backend API for a Todo application built with Go.

## Prerequisites

- Go installed
- PostgreSQL database

## Configuration

The application requires the following environment variables to connect to the database. You can set these in a `.env` file or export them in your shell.

```bash
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_db
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Bangkok
PORT=8080
```

## Running the Application

To start the API server, use the `serve` command:

```bash
go run main.go serve
```