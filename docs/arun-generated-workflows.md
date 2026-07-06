# Generated CI Workflow Follow-up

ARUN generated the workflow definitions below, but the current GitHub token cannot publish files under `.github/workflows/` without the `workflow` scope. Add these workflow files manually or rerun ARUN with a token that includes `workflow` scope.

## `.github/workflows/ci.yml`

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  validate:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: ["1.22"]

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
          cache: true
          cache-dependency-path: server/go.sum

      - name: Download dependencies
        run: |
          cd server
          go mod download

      - name: Run linter
        run: |
          cd server
          go vet ./...

      - name: Run tests
        run: |
          cd server
          go test -v -race -coverprofile=coverage.out ./...

      - name: Build validation
        run: |
          cd server
          go build -v ./...
```

