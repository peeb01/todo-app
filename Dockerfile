FROM public.ecr.aws/docker/library/golang:1.25-alpine3.21 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0

RUN go build -v -a -o main ./main.go

# Runtime
FROM gcr.io/distroless/static-debian12:nonroot


WORKDIR /app
COPY --from=builder /app/main .

USER nonroot:nonroot

ENTRYPOINT ["./main", "serve"]
