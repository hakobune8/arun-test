# Implementation Contract

この文書は、Sprint 1 の product brief を実装可能な技術仕様へ接続するための契約書です。
Backend、Frontend、Deployment、Validation の各担当は、以下のパスと振る舞いを厳守してください。

## 1. Repository Layout
```
.
├── server/          # Go HTTP server & backend logic
│   ├── main.go      # Entry point
│   └── ...
├── client/          # Browser assets (HTML/CSS/JS)
│   ├── index.html   # Primary entrypoint
│   └── ...
├── charts/          # Helm chart for Kubernetes deployment
├── docs/            # Product brief & this contract
├── Dockerfile       # Container build definition
└── README.md        # Project overview
```

## 2. Primary Route & User Path
- **Primary Route**: `GET /`
- **Behavior**: `server/main.go` は `client/index.html` を静的ファイルとして提供し、ブラウザでインベーダーゲームの UI が直接表示されること。
- **Health Check**: `GET /health`
- **Behavior**: `{"status": "ok"}` を返すこと。

## 3. Frontend Contract (`client/`)
- **Entrypoint**: `client/index.html`
- **Required Assets**:
  - Canvas 要素を使用したゲーム描画ロジック。
  - 「新規性のあるインベーダーゲーム」の差別化 mechanic（例：弾の軌道制御や敵の配置パターン）を可視化する UI。
- **Validation**: ブラウザで `/` にアクセスした際、UI が正常にレンダリングされ、操作可能な状態であること。

## 4. Backend Contract (`server/`)
- **Module Path**: `github.com/hakobune8/arun-test/server`
- **Entrypoint**: `server/main.go`
- **Served Routes**:
  - `GET /` -> `client/index.html`
  - `GET /health` -> JSON status
- **Architecture**: HTTP handler と domain logic を分離可能な構成を維持すること。

## 5. Deployment Contract (`charts/` & `Dockerfile`)
- **Dockerfile**: `server/` をビルド対象とし、軽量なベースイメージを使用すること。
- **Helm Chart**: `charts/` 配下に配置。
  - `Service`、`Deployment`、`Selectors`、`Labels`、`Probes`、`Resource Defaults` を含めること。
  - `Ingress` は不要。
- **Validation**: `helm template` でマニフェストが正常に生成されること。

## 6. Validation Commands
以下のコマンドは、fresh checkout 状態で実行可能であり、すべてパスすること。
- `go test ./server/...`
- `docker build -t arun-test .`
- `helm template arun-test ./charts/`
- `curl -s http://localhost:8080/health` (server 起動後)

## 7. Acceptance Criteria Mapping
- **Novelty**: 差別化 mechanic が UI/コードに実装されている。
- **Playful/Simple**: UI が直感的で、操作に迷わない。
- **Production-ready**: Health endpoint、Dockerfile、Helm chart が整備されている。