# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o nyab-backend main.go

# Run stage
FROM alpine:latest
WORKDIR /app

# Needed for Postgres driver
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/nyab-backend .
ENV PORT=8080

EXPOSE 8080
CMD ["./nyab-backend"]
