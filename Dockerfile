FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest

RUN adduser -D mcpuser
USER mcpuser

WORKDIR /app

COPY --from=builder /app/main .

ENV SF_MCP_PORT=9000
EXPOSE 9000

CMD ["./main"]
