# Testing

## Automated validation

```sh
cd server && go test ./...
cd server && go vet ./...
```

## Smoke check

```sh
cd server && go run .
curl http://127.0.0.1:8080/healthz
```

Expected response:

```json
{"status":"ok"}
```
