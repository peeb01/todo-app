FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0

RUN go build -v -a -o main ./main.go

# Runtime
FROM alpine:3.22

WORKDIR /app
COPY --from=builder /app/main .

ENTRYPOINT ["./main", "serve"]