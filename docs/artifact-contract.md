# Artifact Contract: Echoes

## 1. Product Concept Reference
- **Product Name**: Echoes
- **Core Mechanic**: Predictive Reflex - The player must click the *previous* position of a moving target, not the current one.
- **Target User**: Casual players seeking a novel, simple reflex challenge.
- **Differentiation**: Unlike standard "whack-a-mole" or click-the-target games, this requires anticipating the target's past state, creating a unique cognitive loop.

## 2. Primary Route
- **Path**: `GET /`
- **Behavior**: Serves the main game interface.
- **Requirement**: The UI must clearly display the game title "Echoes" and instructions explaining the "Echo" mechanic (e.g., "Click where the target WAS").

## 3. Frontend Implementation
- **Directory**: `client/`
- **Entrypoint**: `client/index.html`
- **Required Assets**:
  - `client/style.css`: Minimalist, high-contrast styling (e.g., dark background, bright target) to emphasize the "void" theme.
  - `client/script.js`: Game logic handling the core loop: target movement, tracking the "previous" position, and click validation.
- **Validation**: The game must be playable immediately upon loading `client/index.html` served by the backend.

## 4. Backend Implementation
- **Directory**: `server/`
- **Module Path**: `github.com/hakobune8/arun-test/server`
- **Entrypoint**: `server/main.go`
- **Served Routes**:
  - `GET /`: Serves `client/index.html` (embedded or static).
  - `GET /health`: Returns `200 OK` with JSON body `{"status": "ok"}`.
- **Logic**:
  - Use Go's `embed.FS` to embed frontend assets for single-binary deployment.
  - Serve static files from the embedded filesystem.

## 5. Deployment & Packaging
- **Dockerfile**: Root level. Multi-stage build: Go build for server, copy client assets, run server.
- **Helm Chart**: `charts/echoes/`
  - `Chart.yaml`: Name `echoes`, version `0.1.0`.
  - `values.yaml`: Default image, port `8080`.
  - `templates/`: `deployment.yaml`, `service.yaml`.
  - **Constraints**: No Ingress. Self-contained chart directory. Service and Deployment must use matching selectors/labels.

## 6. Validation Commands
- **Local Build**: `cd server && go build -o ../server/server .`
- **Run**: `./server/server` (or `go run server/main.go`)
- **Smoke Test**:
  - `curl -s http://localhost:8080/health` (Expect `{"status":"ok"}`)
  - `curl -s http://localhost:8080/` (Expect HTML content containing "Echoes")
- **Tests**: `go test ./...`

## 7. Repository Layout
```text
/
├── .gitignore
├── Dockerfile
├── README.md
├── docs/
│   ├── product-brief.md
│   └── artifact-contract.md
├── server/
│   ├── main.go
│   └── ...
├── client/
│   ├── index.html
│   ├── style.css
│   └── script.js
└── charts/
    └── echoes/
        ├── Chart.yaml
        ├── values.yaml
        └── templates/
```
