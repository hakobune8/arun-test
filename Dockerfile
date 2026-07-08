# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./server

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /server /app/server

# Copy frontend assets (assuming they are in client/ and served by Go)
COPY client/ /app/client/

# Expose port
EXPOSE 8080

# Run
CMD ["/app/server"]
