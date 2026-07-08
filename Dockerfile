# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy Go module files first for layer caching
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy source code
COPY server/ .

# Build the binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server .

# Final stage
FROM alpine:3.21

WORKDIR /app

# Install ca-certificates for HTTPS if needed
RUN apk --no-cache add ca-certificates

# Copy the built binary
COPY --from=builder /server /app/server

# Copy client assets for static serving
COPY client/ /app/client/

# Expose the port the server listens on
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the server
CMD ["/app/server"]
