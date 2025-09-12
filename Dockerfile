## Dockerfile
# Stage 1: Build the Go application

FROM golang:1.25 AS builder

# установка рабочей директории
WORKDIR /app

# копирование go.mod и go.sum для загрузки зависимостей
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

# копирование остального кода приложения
COPY ./app ./


# сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Stage 2: Create a minimal image to run the application
FROM debian:bullseye-slim


WORKDIR /app

# копирование скомпилированного бинарного файла из стадии сборки
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
