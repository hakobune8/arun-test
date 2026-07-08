# sprint-2-plan Recovery Plan

ARUN generated this deterministic planning artifact after the built-in analyst
runtime failed before producing a usable plan.

## Scope

Create a concise product plan and artifact contract for the requested implementation slice.

Key recovered requirements:
- Health endpoint、明確な startup/configuration、focused tests、小さな product API または static asset handler を持つ minimal Go HTTP server。
- Repository goal に合う場合の minimal Web UI または static frontend slice。placeholder だけの画面ではなく、review 可能な primary user path を用意してください。
- Dockerfile と local validation commands。
- ARUN と同じ Kubernetes environment に deploy できる Helm chart と Kubernetes manifests。Chart は charts/<name>/ 配下で Chart.yaml, values.yaml, templates/ を自己完結させ、chart subdirectory の横に charts/values.yaml や charts/Chart.yaml のような孤立 chart 断片を残さないでください。Service、Deployment、selectors、labels、可能な範囲の probes、resource defaults を含めてください。Ingress は不要です。
- Tests、lint または smoke checks、build validation を実行する GitHub Actions CI。

## Implementation Direction

- Treat `docs/product-brief.md` as the product source of truth.
- Treat `docs/artifact-contract.md` as the connection contract for routes,
  frontend assets, backend entrypoints, module path, and validation commands.
- Start with the smallest reviewable Go `net/http` service increment.
- Keep `/healthz` available and covered by tests.
- Add or preserve a lightweight frontend/static response only when it helps the
  repository goal.
- Keep follow-up implementation stages responsible for concrete code changes.

## Validation Expectations

- Run `cd server && go test ./...` when Go sources are present.
- Run `cd server && go vet ./...` when the Go toolchain is available.
- Record smoke-test evidence in repository documentation.
