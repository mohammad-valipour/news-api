# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o news-api .

# Run stage
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/news-api ./news-api

# ایجاد فایل خالی برای godotenv
COPY --from=builder /dev/null .env

EXPOSE 2345

CMD ["./news-api"]