# syntax=docker/dockerfile:1

# Build Go server
FROM golang:1.22-alpine AS server-builder
WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./...

# Final runtime image
FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=server-builder /server /app/server
COPY client/ /app/client/
EXPOSE 8080
CMD ["/app/server"]
