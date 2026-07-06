# Testing

## Automated validation

```sh
go test ./...
go vet ./...
```

## Smoke check

```sh
go run ./server
curl http://127.0.0.1:8080/healthz
```

Expected response:

```json
{"status":"ok"}
```
