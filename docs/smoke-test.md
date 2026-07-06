# Smoke Test

1. Run `go test ./...`.
2. Run `go vet ./...`.
3. Start the service with `go run ./server`.
4. Request `http://127.0.0.1:8080/healthz` and confirm the JSON status is `ok`.
5. Request `/` and confirm the service returns a successful response.
