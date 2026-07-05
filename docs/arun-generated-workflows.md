# Generated CI Workflow Follow-up

ARUN generated the workflow definitions below, but the current GitHub token cannot publish files under `.github/workflows/` without the `workflow` scope. Add these workflow files manually or rerun ARUN with a token that includes `workflow` scope.

## `.github/workflows/ci.yml`

```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23'
          cache: true

      - name: Download dependencies
        run: go mod download

      - name: Vet code
        run: go vet ./...

      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          file: ./coverage.out
        continue-on-error: true

      - name: Build
        run: go build -v ./...

      - name: Smoke test
        run: |
          go build -o /tmp/invader-game .
          /tmp/invader-game &
          sleep 2
          curl -sf http://localhost:8080/healthz || exit 1
          kill %1 2>/dev/null || true
```

