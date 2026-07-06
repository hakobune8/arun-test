# Artifact Contract

この文書は、プロダクトの設計意図（Product Brief）と実装成果物（Implementation Artifacts）を接続する唯一の信頼情報源（Source of Truth）です。

## 1. Product Concept: GravInvader

*   **Core Loop**: インベーダーを倒しつつ、プレイヤーが「重力」を操作して壁や天井を移動する。
*   **Differentiation**: 従来の上下左右移動に加え、`Space`キーで重力方向を反転させることで、敵の弾を回避したり、天井から落下して攻撃したりする。
*   **Target User**: シンプルな操作で爽快感を求めるアーケードゲームファン。

## 2. Primary Route & Serving

*   **Entry Point**: `GET /`
    *   **Response**: `client/index.html` を Content-Type `text/html` で返す。
*   **Static Assets**: `GET /assets/*`
    *   **Response**: `client/assets/` 配下のファイル（CSS, JS, Images）を返す。
*   **Health Check**: `GET /health`
    *   **Response**: `200 OK` with JSON `{"status": "ok"}`.

## 3. Frontend Contract

*   **Entrypoint**: `client/index.html`
*   **Assets Directory**: `client/assets/`
    *   `style.css`: ゲーム画面のスタイリング。
    *   `game.js`: Canvasベースのゲームロジック（Canvas API使用）。
*   **UI Requirements**:
    *   ゲームタイトル「GravInvader」の表示。
    *   スコア表示。
    *   「重力反転」操作のヒント表示。

## 4. Backend Contract

*   **Module Path**: `github.com/hakobune8/arun-test`
*   **Entrypoint**: `server/main.go`
*   **Dependencies**:
    *   `net/http` (Standard library)
    *   `embed` (Static assets embedding)
*   **Configuration**:
    *   Port: `8080` (Env var `PORT` で上書き可能)

## 5. Deployment Contract

*   **Dockerfile**: `Dockerfile` (Multi-stage build)
*   **Helm Chart**: `charts/arun-test/`
    *   **Service**: ClusterIP (Port 80 -> 8080)
    *   **Deployment**: Replicas 1, Resource limits defined.
*   **Kubernetes Manifests**: `k8s/` (Optional direct manifests)

## 6. Validation Commands

*   **Local Build**: `go build ./...`
*   **Local Run**: `go run ./server/main.go`
*   **Smoke Test**: `curl -f http://localhost:8080/health`
*   **Docker Build**: `docker build -t gravinvader:latest .`
*   **Helm Lint**: `helm lint charts/arun-test/`

## 7. QA & Acceptance Criteria

*   [ ] `GET /` でゲーム画面（Canvas）が表示される。
*   [ ] `GET /health` で正常なステータスコードが返る。
*   [ ] ゲーム内で重力反転（Spaceキー）が動作する。
*   [ ] Docker イメージが正常にビルドされる。
*   [ ] Helm Chart が lint を通る。
