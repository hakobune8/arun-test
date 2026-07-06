# Artifact Contract: Product Brief Md

This contract is the implementation source of truth for how generated artifacts connect. Product intent lives in `docs/product-brief.md`; this file defines the route, file, module, and validation expectations that backend, frontend, QA, Docker, Helm, and documentation work must preserve.

## Primary route

- `/` serves the primary browser experience.
- `/healthz` returns a JSON health response.

## Frontend

- Directory: `client/`.
- Package file: `client/package.json`.
- Entrypoint: `client/index.html`.
- Required local assets: `client/styles.css` and `client/src/main.js`.
- HTML must not reference local CSS or JavaScript files that are absent from the repository.

## Backend

- Language: Go.
- Module path: `github.com/hakobune8/arun-test/server`.
- Module file: `server/go.mod`.
- Entrypoint: `server/main.go`.
- The Go server must serve `/` from `client/index.html` and must serve every local CSS or JavaScript file referenced by that HTML.

## Deployment

- Docker and Kubernetes artifacts must package the same backend and frontend paths defined above.
- Helm chart templates must expose the application service and health checks without introducing a separate product path.

## Validation

- `cd server && go test ./...` passes when Go tooling is available.
- `cd server && go vet ./...` passes when Go tooling is available.
- Frontend smoke validation confirms `/` returns HTML and that referenced local CSS/JS assets exist and are served.
- QA must treat any mismatch between this contract and repository artifacts as release-blocking.
