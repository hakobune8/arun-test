# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files first for caching
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy source and build
COPY server/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# Runtime stage
FROM alpine:3.21

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

# Run
CMD ["./server"]
