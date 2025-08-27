FROM golang:1.25 AS builder

WORKDIR /app

COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

COPY ./app ./


RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
