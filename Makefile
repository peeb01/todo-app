BINARY_NAME=todo-app

.PHONY: all build run clean test
.PHONY: all build run clean test swag

all: build

build:
	go build -o $(BINARY_NAME) main.go

run:
	go run main.go serve

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe

test:
	go test ./...
