# ---- Build stage ----
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /newsletter-api ./cmd

# ---- Final stage ----
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /newsletter-api .

EXPOSE 8080

CMD ["./newsletter-api"]