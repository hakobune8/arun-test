# Artifact Contract

この文書は、Sprint 1 の実装におけるファイル構成、ルーティング、パッケージング、検証手順の契約（Contract）を定義します。

## 1. リポジトリレイアウト

関心を分離し、拡張性を確保するため、以下のディレクトリ構成を採用します。

```
.
├── server/          # Go HTTP サーバーとバックエンドロジック
│   ├── main.go      # エントリポイント
│   ├── handlers/    # HTTP ハンドラー
│   └── game/        # ドメインロジック（将来の分離用）
├── client/          # ブラウザフロントエンド
│   ├── index.html   # ゲーム画面のエントリポイント
│   ├── style.css    # スタイル定義
│   └── script.js    # クライアントロジック
├── charts/          # Helm チャート（Kubernetes 展開用）
│   └── arun-test/
├── docs/            # 製品仕様と検証ドキュメント
│   ├── product-brief.md
│   └── artifact-contract.md
├── Dockerfile       # コンテナビルド定義
├── go.mod           # Go モジュール定義（ルートまたは server/配下）
└── README.md        # プロジェクト概要
```

## 2. Primary Route（主要ルート）

- **URL**: `GET /`
- **説明**: ゲームのメイン画面（`client/index.html`）を提供します。
- **期待される動作**: ブラウザでアクセスすると、インベーダーゲームの UI が表示され、操作可能であること。

## 3. Frontend 契約

- **ディレクトリ**: `client/`
- **エントリポイント**: `client/index.html`
- **必須アセット**:
  - `client/style.css`: ゲームの視覚的スタイル。
  - `client/script.js`: ゲームのクライアント側ロジック（描画、入力処理）。
- **要件**: 
  - 外部 CDN 依存を最小限にし、ローカルアセットで動作すること。
  - Canvas または DOM を使用したゲーム描画が初期状態で表示されること。

## 4. Backend 契約

- **モジュールパス**: `github.com/hakobune8/arun-test/server` （またはローカルビルド用として `server`）
- **エントリポイント**: `server/main.go`
- **ロジック分離**: 
  - HTTP ハンドラーは `server/handlers/` に配置。
  - ゲームのドメインロジックは将来 `server/game/` へ分離可能にする。
- **要件**:
  - `client/` ディレクトリを静的ファイルとして提供すること。
  - 外部サービスなしでローカルで起動・動作すること。

## 5. Served Routes（提供されるルート）

| Method | Path           | 説明                          | 提供ファイル/レスポンス      |
|--------|----------------|-------------------------------|-----------------------------|
| GET    | `/`            | ゲームメイン画面              | `client/index.html`         |
| GET    | `/style.css`   | スタイルシート                | `client/style.css`          |
| GET    | `/script.js`   | スクリプトファイル            | `client/script.js`          |
| GET    | `/health`      | ヘルスチェック                | `200 OK`                    |

## 6. Docker / Helm パッケージング期待値

- **Dockerfile**:
  - ルートディレクトリまたは `server/` に配置。
  - Multi-stage build を推奨（ビルド環境とランタイム環境の分離）。
  - 最終イメージは `server/main.go` を実行するものとする。
- **Helm Chart**:
  - ディレクトリ: `charts/arun-test/`
  - 必須リソース: `Deployment`, `Service`。
  - ラベルとセレクターの一貫性を保つこと。
  - Ingress は含めない（Sprint 1 の範囲外）。

## 7. 検証コマンド（Validation Commands）

以下のコマンドは、実装の正しさを確認するために使用します。

### ビルド検証
```bash
# Go バックエンドのビルド
go build ./server/...

# クライアントアセットの確認
ls -la client/
```

### ローカル実行
```bash
# サーバーの起動
go run ./server/main.go

# ブラウザでアクセス
open http://localhost:8080
```

### コンテナ化検証
```bash
# Docker イメージのビルド
docker build -t arun-test:latest .

# コンテナの実行
docker run -p 8080:8080 arun-test:latest
```

### Helm 検証
```bash
# Helm チャートの lint
helm lint charts/arun-test/
```

## 8. 変更対象ファイル一覧

Sprint 1 で新規作成または更新が予定されているファイル:

- `docs/product-brief.md` (新規)
- `docs/artifact-contract.md` (新規/更新)
- `server/main.go` (新規)
- `server/handlers/` (新規)
- `client/index.html` (新規)
- `client/style.css` (新規)
- `client/script.js` (新規)
- `Dockerfile` (新規)
- `charts/arun-test/` (新規)
- `README.md` (新規/更新)
