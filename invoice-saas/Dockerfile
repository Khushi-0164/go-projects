FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /invoice-saas ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /invoice-saas .
COPY db/migrations ./db/migrations

EXPOSE 8080

CMD ["./invoice-saas"]