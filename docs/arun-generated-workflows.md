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
  validate:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: [ '1.21', '1.22' ]
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
          cache: true

      - name: Download dependencies
        run: go mod download

      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...

      - name: Upload coverage
        uses: codecov/codecov-action@v4
        with:
          file: ./coverage.out
          fail_ci_if_error: false

      - name: Build
        run: go build -v ./...

      - name: Lint
        uses: golangci/golangci-lint-action@v4
        with:
          version: latest
          args: --timeout=5m
```

