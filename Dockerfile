# Build stage
FROM golang:1.21-alpine3.17 AS builder
WORKDIR /app

# Copy only the go mod and sum files to download dependencies
COPY go.mod go.sum ./
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080
CMD ["/app/main"]