# Сборка
FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/bot

# Запуск
FROM alpine
WORKDIR /app/
COPY --from=builder /app/bot .
CMD ["./bot"]
