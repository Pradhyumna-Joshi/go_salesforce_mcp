# Build stage
FROM golang:1.26.1-alpine3.23 AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Run stage
FROM alpine:latest

RUN apk update && apk upgrade && rm -rf /var/cache/apk/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 9000

CMD ["./main"]