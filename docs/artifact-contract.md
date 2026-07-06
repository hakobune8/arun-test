# Artifact Contract

この文書は、`docs/product-brief.md` で定義されたプロダクトコンセプトを具体的な実装成果物に接続する契約書です。
バックエンド、フロントエンド、デプロイメントの各アセットは、以下のパスと振る舞いに厳密に従って実装・検証されます。

## 1. Primary Route
- **パス**: `/`
- **振る舞い**: ゲームのメイン画面（Canvas/UI）をレンダリングする HTML ページを返す。
- **期待値**: ブラウザでアクセスした際、タイトル、操作説明、ゲーム開始ボタンが表示される。

## 2. Frontend Directory & Entrypoint
- **ディレクトリ**: `client/`
- **エントリーポイント**: `client/index.html`
- **必須ローカルアセット**:
  - `client/style.css`: ゲーム画面のレイアウト、ポップな配色、レスポンシブ対応を定義。
  - `client/app.js`: ゲームループ、入力処理、描画ロジック、バックエンドAPIとの通信を定義。
- **期待値**: `index.html` は `style.css` と `app.js` を正しく読み込み、ブラウザで直接開いた際にも基本的なUIが確認できること。

## 3. Backend Module Path & Entrypoint
- **モジュールパス**: `server/`
- **エントリーポイント**: `server/main.go`
- **必須ロジック**:
  - HTTP サーバーの起動と設定。
  - `/health` エンドポイントの実装。
  - `/` での静的ファイル提供（`client/` 配下）。
  - ゲームの状態管理用 API（例: `/api/game`）のstubまたは実装。
- **期待値**: `go run main.go` でサーバーが起動し、`/health` が `200 OK` を返すこと。

## 4. Served Routes
| パス | 提供元 | 内容 |
|---|---|---|
| `/` | `server/` | `client/index.html` を提供 |
| `/style.css` | `server/` | `client/style.css` を提供 |
| `/app.js` | `server/` | `client/app.js` を提供 |
| `/health` | `server/` | `{"status": "ok"}` をJSONで返す |
| `/api/game` | `server/` | ゲーム状態の取得/更新用エンドポイント（stub含む） |

## 5. Docker/Helm Packaging Expectations
- **Dockerfile**: リポジトリルートまたは `server/` 直下に配置。マルチステージビルドで Goバイナリを生成し、軽量なランタイムイメージを作成。
- **Helm Chart**: `charts/` ディレクトリに配置。
  - `values.yaml`: リソース制限、レプリカ数、環境変数を定義。
  - `templates/deployment.yaml`: Deployment、Service、Selector、Label、Liveness/Readiness Probe、リソースデフォルトを含む。
  - `templates/service.yaml`: ClusterIPまたはNodePortのService定義。
- **期待値**: `helm template .` で有効なKubernetesマニフェストが出力されること。`docker build` でイメージが正常にビルドされること。

## 6. Validation Commands
以下のコマンドは、実装が契約を満たしていることをQAが確認するために使用します。

```bash
# 1. Goサーバーのビルドと起動確認
cd server && go build -o server . && ./server &
# 2. Healthエンドポイントの検証
curl -s http://localhost:8080/health | grep -q "ok"
# 3. Primary Routeの検証
curl -s http://localhost:8080/ | grep -q "<html"
# 4. Frontendアセットの検証
curl -s http://localhost:8080/style.css | grep -q "body"
curl -s http://localhost:8080/app.js | grep -q "game"
# 5. Helmテンプレートの検証
helm template test-release charts/ | grep -q "kind: Deployment"
helm template test-release charts/ | grep -q "kind: Service"
# 6. Goテストの実行
cd server && go test ./...
```

## 7. Repository Layout
```
.
├── server/          # Go HTTPサーバーとバックエンドロジック
│   ├── main.go
│   └── ...
├── client/          # ブラウザアセット（HTML/CSS/JS）
│   ├── index.html
│   ├── style.css
│   └── app.js
├── charts/          # HelmチャートとKubernetesマニフェスト
│   ├── Chart.yaml
│   ├── values.yaml
│   └── templates/
├── docs/            # プロダクトと検証ドキュメント
│   ├── product-brief.md
│   └── artifact-contract.md
├── Dockerfile       # コンテナビルド定義
├── README.md        # プロジェクト概要とクイックスタート
└── .github/         # CI/CDワークフロー
```

この契約書は、実装フェーズにおけるすべてのファイル配置、パス、振る舞いの基準となります。