# Implementation Contract

この文書は、`docs/product-brief.md` で定義されたプロダクトコンセプトを技術的な実装に接続するための契約です。

## 1. Source of Truth

- **Product Brief**: `docs/product-brief.md`
- **Concept**: リズムに合わせた射撃が核心となる「リズムインベーダー」。
- **Differentiating Mechanic**: ビジュアルまたはオーディオのビートに合わせて敵を撃つと「パーフェクト」判定となり、スコアボーナスや視覚的フィードバックが発生する。

## 2. Repository Layout

関心を分離し、拡張性を確保するため以下のディレクトリ構成を採用します。

```
.
├── server/          # Go HTTP サーバーとバックエンドロジック
│   ├── main.go      # エントリポイント
│   ├── go.mod
│   └── ...
├── client/          # ブラウザ用フロントエンドアセット
│   ├── index.html   # ゲームUIのエントリポイント
│   ├── style.css    # スタイル定義
│   └── app.js       # ゲームロジック（Canvas API使用）
├── charts/          # Helmチャート（Kubernetesデプロイメント用）
│   └── ...
├── docs/            # プロダクトおよび検証ドキュメント
│   ├── product-brief.md
│   └── artifact-contract.md
├── Dockerfile       # コンテナビルド定義
├── Makefile         # ローカル検証コマンドのエイリアス
└── README.md        # プロジェクト概要
```

## 3. Backend Contract

### Module & Entrypoint
- **Module Path**: `github.com/hakobune8/arun-test/server`
- **Entrypoint**: `server/main.go`

### Routes & Behavior
1. **`GET /`**
   - `client/index.html` を serve する。
   - ゲームUIが直接表示されるプライマリルート。
2. **`GET /health`**
   - `200 OK` を返す。Kubernetesのliveness/readiness probe用。
3. **`GET /static/*`**
   - `client/` 内の CSS/JS ファイルを serve する。

### Architecture Note
- 現在のSprint 1では、HTTPハンドラとドメインロジックを `server/` 内で管理するが、将来の拡張（WebSocket通信、スコアDB連携など）を見据え、`server/handler/` と `server/game/` への分離を可能にする構造を維持する。

## 4. Frontend Contract

### Entrypoint & Assets
- **HTML**: `client/index.html`
- **CSS**: `client/style.css`
- **JS**: `client/app.js`

### Implementation Requirements
1. **Canvas Rendering**: `app.js` は HTML5 Canvas APIを使用してゲームループを実装する。
2. **Mechanic Implementation**: 
   - 「リズムインベーダー」の差別化 mechanic（ビート同期射撃）を `app.js` で実装する。
   - 例：画面上にビートインジケーターを表示し、敵が特定の位置に来たタイミングで射撃すると視覚的エフェクトが発生するロジック。
3. **Responsiveness**: 基本的なレスポンシブ対応（または固定アスペクト比のコンテナ）を実装し、ブラウザで正常に表示できることを確認する。

## 5. Deployment Contract

### Docker
- **Dockerfile**: リポジトリルートまたは `server/` に配置。
- **Multi-stage Build**: Goバイナリビルドと静的ファイルの結合を含むマルチステージビルドを採用する。
- **Image**: `arun-test:latest`

### Kubernetes / Helm
- **Charts**: `charts/` ディレクトリにHelmチャートを配置。
- **Components**:
  - `Deployment`: リソース制限（requests/limits）とプローブ（liveness/readiness）を含む。
  - `Service`: ClusterIPまたはNodePort（環境による）。
  - **Ingress**: 不要（Sprint 1時点）。

## 6. Validation Commands

以下のコマンドで検証可能であることを確認する。

1. **Build & Test**
   ```bash
   go build ./server/...
   go test ./server/...
   ```
2. **Local Run**
   ```bash
   go run ./server/main.go
   # http://localhost:8080 にアクセスしてゲームUIが表示されることを確認
   ```
3. **Docker Build**
   ```bash
   docker build -t arun-test .
   docker run -p 8080:8080 arun-test
   ```
4. **Helm Template**
   ```bash
   helm template arun-test ./charts/
   ```

## 7. Acceptance Criteria (Technical)

- [ ] `server/main.go` が起動し、`/` で `client/index.html` が返される。
- [ ] `client/app.js` がゲームループを実行し、Canvasに描画する。
- [ ] 「リズムインベーダー」の核心 mechanic がコード内に実装されている。
- [ ] Dockerfile が正常にビルドできる。
- [ ] Helm chart が `helm template` で正常にレンダリングされる。
