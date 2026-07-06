# Artifact Contract

## Overview
This document serves as the implementation contract connecting the product brief to the codebase. It defines the file layout, primary routes, entry points, and validation commands.

## File Layout
- `server/`: Go HTTP server and backend logic.
- `client/`: Browser assets (HTML, CSS, JS) for the game UI.
- `docs/`: Product brief, artifact contract, and validation notes.
- `charts/`: Helm charts for Kubernetes deployment.

## Primary Route
- `GET /`: Serves the main game UI (`client/index.html`).
- `GET /health`: Health check endpoint.

## Backend
- **Module Path:** `github.com/hakobune8/arun-test`
- **Entrypoint:** `server/main.go`
- **Logic:** `server/handler.go` (HTTP handlers), `server/game.go` (Game logic)

## Frontend
- **Entrypoint:** `client/index.html`
- **Assets:** `client/style.css`, `client/app.js`

## Deployment
- **Dockerfile:** Root directory
- **Helm Chart:** `charts/arun-test/`

## Validation Commands
- `go build ./...`
- `go test ./...`
- `docker build -t arun-test .`
- `helm lint charts/`
