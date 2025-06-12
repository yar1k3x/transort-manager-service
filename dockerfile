# Этап сборки
FROM golang:1.24.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ВАЖНО: Статическая сборка, без glibc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o transport-management-service main.go

# Финальный образ — пустой (без glibc, минимальный размер)
FROM scratch

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/transport-management-service .
COPY --from=builder /app/service ./service
COPY --from=builder /app/.env ./.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Запускаем
ENTRYPOINT ["./transport-management-service"]
