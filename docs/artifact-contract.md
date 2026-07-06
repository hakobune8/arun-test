# Artifact Contract

この文書は、`docs/product-brief.md` の製品要件を実装成果物へ接続する source of truth です。QA、レビュー、実装のいずれもこの contract に従って検証してください。

## 1. Product Brief 参照

- **Product Brief**: `docs/product-brief.md`
- **Concept**: 新規性のあるインベーダーゲーム（ポップでシンプル、ブラウザで動作）
- **Primary User Path**: ブラウザで `/` にアクセス → ゲーム画面が表示 → 操作可能

## 2. Primary Route

| Route | 説明 | 実装ファイル |
|-------|------|-------------|
| `/` | メインゲーム画面（HTML + 埋め込み CSS/JS） | `server/main.go` の `handleIndex` |
| `/healthz` | ヘルスチェックエンドポイント | `server/main.go` の `handleHealthz` |
| `/assets/style.css` | ゲームスタイルシート | `server/main.go` の `handleAssets` |
| `/assets/game.js` | ゲームロジック | `server/main.go` の `handleAssets` |

## 3. Frontend

| 項目 | 詳細 |
|------|------|
| **ディレクトリ** | `client/` |
| **エントリーポイント** | `client/index.html` |
| **CSS** | `client/assets/style.css` |
| **JS** | `client/assets/game.js` |
| **Go からの提供** | `server/main.go` で `embed.FS` を使用し、`client/` 配下の全アセットを `/assets/` にマウント |

## 4. Backend

| 項目 | 詳細 |
|------|------|
| **Go モジュールパス** | `github.com/hakobune8/arun-test` |
| **エントリーポイント** | `server/main.go` |
| **依存関係** | 標準ライブラリのみ（`net/http`, `embed`, `log`, `os`） |
| **構成** | 環境変数 `PORT`（デフォルト: `8080`） |

## 5. Deployment

| 項目 | 詳細 |
|------|------|
| **Dockerfile** | `Dockerfile`（multi-stage build） |
| **Helm Chart** | `charts/arun-test/` |
| **K8s Manifests** | `k8s/`（Deployment, Service） |
| **Service Type** | ClusterIP（Ingress なし） |
| **Probes** | Liveness: `/healthz`, Readiness: `/healthz` |
| **Resource Defaults** | requests: 32Mi/50m, limits: 128Mi/200m |

## 6. Validation Commands

| チェック | コマンド |
|---------|---------|
| **ビルド** | `go build ./...` |
| **テスト** | `go test ./...` |
| **ローカル実行** | `PORT=8080 go run server/main.go` |
| **ヘルスチェック** | `curl http://localhost:8080/healthz` |
| **メイン画面** | `curl http://localhost:8080/` |
| **アセット提供** | `curl http://localhost:8080/assets/style.css` |
| **Docker ビルド** | `docker build -t arun-test .` |
| **Helm lint** | `helm lint charts/arun-test/` |

## 7. QA 検証ポイント

1. `/` にアクセスするとゲームタイトルと操作説明が表示される
2. `/healthz` が `200 OK` を返す
3. `/assets/style.css` と `/assets/game.js` が正常に提供される
4. ゲーム画面でキーボード操作（矢印キー + スペース）でプレイヤー移動と弾発射が可能
5. 敵が画面を移動し、弾が敵に当たると消滅する
6. Docker イメージが正常にビルドされる
7. Helm chart が `helm lint` をパスする

## 8. Concept Drift 防止

- 製品名は **「Pop Invaders」** のみを使用
- 差別化 mechanic: **「ポップなビジュアル + シンプルな操作 + リズム感ある敵の動き」**
- 複数の product brief、別名、矛盾した mechanic を作成しない
- 変更時はこの contract を更新し、product-brief.md と同期する
