# Arun Test

This repository started empty. ARUN generated a minimal static browser game with a gravity-lane mechanic so an implementation-heavy scrum workflow can produce reviewable code, documentation, and validation artifacts without GitHub API calls.

## Repository layout

- `server/` contains the Go HTTP entrypoint.
- `client/` contains the browser UI served from `/`.
- `charts/` and `k8s/` contain deployment artifacts when present.
- `docs/` contains product and validation notes.

## Features

- Keyboard controls with ArrowLeft, ArrowRight, and Space.
- Space flips the defender between floor and ceiling gravity lanes.
- Score display that increments only when the defender is horizontally aligned and on the same gravity lane as the invader.
- Lives tracking that decrements when an invader reaches the bottom of the arena.
- Restart behavior that resets score, lives, player position, and invader position.

## Run

Run the Go server with `cd server && go run .` and open `http://127.0.0.1:8080/`, or open `client/index.html` directly for a static browser review.

## Validate

```sh
cd server && go test ./...
cd server && go vet ./...
npm --prefix client test
npm --prefix client run build
for chart in charts/*; do [ -f "$chart/Chart.yaml" ] && helm lint "$chart" && helm template arun-validation "$chart" >/tmp/arun-validation.yaml; done
```

The client scripts use `node --check` from `client/package.json` and do not require package installation. Run the Helm commands after the chart is present.
