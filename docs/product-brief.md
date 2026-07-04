# Product Brief: Kubernetes

## Concept

Kubernetes is a compact browser invader game built around a gravity-lane flip mechanic. The player moves left and right, then uses Space to shift between the floor and ceiling lanes. Scoring requires both horizontal alignment and matching the invader lane, so the differentiating mechanic is present in the implemented UI and source code rather than only in documentation.

## Target User

- Players who want a short arcade loop with one clear twist.
- Reviewers who need a fresh-checkout slice that runs without external services.

## Acceptance Criteria

- The visible title, README H1, and this product brief use the same product name.
- The primary route `/` serves the browser game when run through the Go server.
- Space changes the gravity lane between Floor and Ceiling.
- A score is awarded only when the defender is aligned with the invader and on the same lane.
- The Docker runtime image includes the static assets required for `/` to serve the same UI as local `go run .`.

## Non-Goals

- Multiplayer.
- External score services.
- Complex level progression.

## Source Request

Sprint 3 documentation: update README or docs with a product-centered overview, primary user walkthrough, acceptance criteria status, local run, test, Docker, Helm/Kubernetes deploy, rollback or operations notes, and reviewer guidance. README H1 must be the product name or repository name, not a deployment topic such as Kubernetes. Explain what was built and how it behaves before listing commands. Keep README concise, move detailed procedures into focused docs, remove duplicated or stale instructions, remove copied parent-task prompt text, remove duplicate product brief files including product/design.md when it competes with docs/product-brief.md, and remove links to non-existent sprint reports or operations docs. Documentation must describe the implemented product, not an aspirational or alternate concept. Artifact hygiene: do not copy the full parent task, prompt text, run workspace contents, compiled binaries, or ARUN-generated archive artifacts into repository documentation or product files. Summarize only the relevant product requirements and keep generated reports concise.
