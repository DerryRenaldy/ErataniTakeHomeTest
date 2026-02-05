FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for go modules)
RUN apk add --no-cache git

# Copy go mod file
COPY go.mod ./

# Copy source code first (needed for go mod tidy to analyze dependencies)
COPY . .

# Generate go.sum and download dependencies
RUN go mod tidy && go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/server .

# Runtime stage
FROM alpine:latest

# Install necessary packages and dbmate
RUN apk add --no-cache curl ca-certificates && \
    curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64 && \
    chmod +x /usr/local/bin/dbmate

# Create non-root user
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# Set working directory
WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/bin/server ./server

# Copy environment file
COPY --from=builder /app/.env.docker ./.env

# Copy migrations and seed data
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seed-data ./seed-data

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose application port
EXPOSE 8080

# Entrypoint
CMD ["./server"]
