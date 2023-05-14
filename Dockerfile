# Build stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o backend cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /app/configs configs/

EXPOSE 8080
ENTRYPOINT [ "/app/backend" ]
