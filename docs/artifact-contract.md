# Artifact Contract

この文書は、`docs/product-brief.md` で定義されたプロダクト概念を実際のコード、アセット、デプロイメント構成に接続するための単一信頼源（Source of Truth）です。

## 1. リポジトリレイアウト

関心を分離し、レビューと検証を容易にするため、以下のディレクトリ構成を採用します。

- `server/` : Go HTTP エントリポイント、バックエンドロジック、`go.mod`
- `client/` : ブラウザアセット（HTML, CSS, JS）、ゲームUI
- `docs/`   : プロダクトブリーフ、検証ノート、この契約書
- `charts/` : Helm Chart（Kubernetes デプロイメント用）
- `Dockerfile` : サーバーコンテナ化用
- `README.md` : プロダクト概要、セットアップ、検証コマンド

## 2. Primary Route & Served Routes

- `GET /` : ゲームUI（`client/index.html`）と関連アセットをサーブする。これがプライマリーユーザーパス。
- `GET /health` : システムヘルスチェック。`{"status": "ok"}` を返す。
- `GET /assets/*` : `client/` 内の静的アセット（CSS, JS, 画像）をサーブする。

## 3. Frontend Directory & Entrypoint

- **ディレクトリ**: `client/`
- **エントリポイント**: `client/index.html`
- **必須ローカルアセット**:
  - `client/style.css` : ゲームUIのスタイリング
  - `client/app.js` : ゲームロジックとレンダリング（Canvas API使用）
- **差別化メカニクス実装場所**: `client/app.js` 内で「重力反転（Gravity Shift）」メカニクスを実装し、プレイヤーが敵弾を避けるだけでなく、敵の配置を物理的に反転させて攻撃できることを確認する。

## 4. Backend Module Path & Entrypoint

- **モジュールパス**: `github.com/hakobune8/arun-test/server`
- **エントリポイント**: `server/main.go`
- **責務**:
  - HTTP サーバーの起動とルーティング設定
  - `/health` エンドポイントの実装
  - `client/` 内の静的アセットをサーブするファイルサーバーの設定
  - 環境変数からのポート設定（`PORT`、デフォルト: `8080`）

## 5. Docker/Helm Packaging Expectations

- **Dockerfile**:
  - 多段ビルド（Multi-stage build）を使用し、最終イメージは `server/` のバイナリのみを含む。
  - 非rootユーザーで実行し、セキュリティベストプラクティスに従う。
- **Helm Chart (`charts/`)**:
  - `Chart.yaml`, `values.yaml`
  - `templates/deployment.yaml` : Deployment, Replicas, Labels, Selectors, Probes (liveness/readiness), Resource Defaults
  - `templates/service.yaml` : Service (ClusterIP), Ports, Selectors
  - Ingress は含まない（Kubernetes内部アクセスのみを想定）

## 6. Validation Commands

実装完了後、以下のコマンドで検証を行う。

```bash
# Go ビルドとテスト
$ go build ./server/...
$ go test ./server/...

# Docker イメージビルド
$ docker build -t arun-test:latest .

# Helm Lint
$ helm lint charts/

# ローカル実行とSmoke Test
$ go run ./server/main.go &
$ curl -s http://localhost:8080/health
$ curl -s http://localhost:8080/ | head -n 1
```

## 7. 変更が期待されるファイル

- `docs/artifact-contract.md` (このファイル)
- `docs/product-brief.md`
- `server/main.go`
- `server/go.mod`
- `client/index.html`
- `client/style.css`
- `client/app.js`
- `Dockerfile`
- `charts/` 内の全ファイル
- `README.md`

---
この契約書は、Sprint 1 の実装フェーズにおいて、すべてのコード変更、ドキュメント更新、テスト追加の基準となります。プロダクトブリーフと矛盾する変更は行わない。