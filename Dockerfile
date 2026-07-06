# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy Go module files and download dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy source code and build
COPY server/ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /server ./...

# Final stage
FROM alpine:3.21
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /server /app/server

# Copy frontend assets
COPY client/ /app/client

# Expose port
EXPOSE 8080

# Run the server
CMD ["/app/server"]
