# Artifact Contract

この文書は、`docs/product-brief.md` の製品コンセプトを具体的な実装ファイル、ルート、検証コマンドに接続する契約書です。実装・レビュー・QAはこの契約に従って進めます。

## 1. Primary Route & User Path
- **Primary Route**: `GET /`
- **User Experience**: ブラウザで `/` にアクセスすると、ゲームタイトル画面と開始ボタンが表示される。クリックするとゲームキャンバスが描画され、操作可能になる。
- **Health Endpoint**: `GET /health` (JSON: `{"status":"ok"}`)

## 2. Frontend Layout & Entrypoint
- **Directory**: `client/`
- **Entrypoint**: `client/index.html`
- **Required Local Assets**:
  - `client/style.css` (ゲームUI、キャンバス、メニューのスタイリング)
  - `client/game.js` (Canvas APIを使用したゲームループ、重力井戸メカニクス、入力処理、スコア管理)
- **Served Paths**: `/` は `client/index.html` を返し、`/assets/` は `client/` 配下の静的ファイルを提供する。

## 3. Backend Module Path & Entrypoint
- **Module Path**: `github.com/hakobune8/arun-test`
- **Entrypoint**: `server/main.go`
- **Served Routes**:
  - `GET /` -> `client/index.html` を Content-Type `text/html` で提供
  - `GET /health` -> JSON 200 OK
  - `GET /assets/*` -> `client/` 配下のファイルを提供
- **Architecture Note**: 小さな vertical slice では `server/main.go` にハンドラと静的ファイル提供をまとめるが、ドメインロジックは将来 `server/internal/` に分離可能にする。

## 4. Docker / Helm Packaging Expectations
- **Dockerfile**: ルート直下に配置。Go ビルドステージと静的ファイル提供を行う。
- **Helm Chart**: `charts/arun-test/` 配下に配置。
  - `Chart.yaml`, `values.yaml`, `templates/deployment.yaml`, `templates/service.yaml`
  - Service, Deployment, selectors, labels, probes, resource defaults を含む。
  - Ingress は含まない。
  - 孤立した断片 (`charts/values.yaml` など) を残さない。

## 5. Validation Commands
- **Local Build & Run**:
  ```bash
  go build -o bin/server ./server/
  ./bin/server
  ```
- **Smoke Test**:
  ```bash
  curl -s http://localhost:8080/health | jq .
  curl -s http://localhost:8080/ | head -n 5
  ```
- **Tests**:
  ```bash
  go test ./...
  ```
- **Helm Lint & Template**:
  ```bash
  helm lint charts/arun-test/
  helm template arun-test charts/arun-test/ | kubectl apply --dry-run=client -f -
  ```

## 6. File Change Summary
- `docs/product-brief.md` (新規)
- `docs/artifact-contract.md` (新規)
- `server/main.go` (新規)
- `client/index.html`, `client/style.css`, `client/game.js` (新規)
- `Dockerfile` (新規)
- `charts/arun-test/Chart.yaml`, `values.yaml`, `templates/*.yaml` (新規)
- `README.md` (新規)
- `.github/workflows/ci.yml` (新規)
