# Chrono Invaders

**Chrono Invaders** は、レトロなインベーダーゲームに「タイムリワインド」機能を搭載した、新規性のあるアーケードゲームです。

## 概要

このプロジェクトは、以下の価値を提供します：
- **新規性**: 敵を倒すだけでなく、過去に戻って戦略をやり直す「タイムリワインド」 mechanic。
- **シンプル**: 直感的な操作と、レトロなビジュアル。
- **Production-Ready**: Go による堅牢なバックエンド、Docker/Kubernetes 対応のデプロイメント。

## 主な機能 (Sprint 1-3)

- [x] タイムリワインド mechanic の実装
- [x] Go HTTP サーバーと静的アセット配信
- [x] Docker コンテナ化
- [x] Kubernetes (Helm) 対応デプロイメント
- [x] 基本的な Smoke Tests

## 開発者向けガイド

### ローカル実行

Go 環境が必要です。

```bash
# ビルドと実行
go build -o main ./server/cmd/main.go
./main

# または直接実行
go run ./server/cmd/main.go
```

ブラウザで `http://localhost:8080` にアクセスしてください。

### Docker

```bash
docker build -t chrono-invaders .
docker run -p 8080:8080 chrono-invaders
```

## ドキュメント

- [Artifact Contract](docs/artifact-contract.md): 成果物の接続仕様
- [Deployment Guide](docs/deployment.md): Kubernetes へのデプロイ手順
- [QA Report](docs/qa-report.md): 検証結果と known limitations

## 構造

```
.
├── server/       # Go backend & HTTP server
├── client/       # Frontend assets (HTML/CSS/JS)
├── charts/       # Helm charts
├── docs/         # Product & operational docs
└── README.md     # This file
```

## ライセンス

MIT License
