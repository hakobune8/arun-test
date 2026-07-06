# 🚀 StarHopper — 新規性インベーダーゲーム

## 概要

**StarHopper** は、従来のインベーダーゲームに「重力波シールド」メカニクスを追加した、シンプルかつ戦略的なシューティングゲームです。

プレイヤーは宇宙船を操作し、波状に襲いかかる敵インベーダーを撃破します。従来のインベーダーゲームとの差別化ポイントとして、**重力波シールド**（Gravity Wave Shield）を搭載。シールドを張ると敵弾を跳ね返せますが、シールド使用中は自機が固定されるため、タイミングと位置取りが勝敗を分けます。

## 主な機能（Sprint 1 実装済み）

- [x] ゲーム画面のレンダリング（HTML5 Canvas）
- [x] 自機操作（左右移動、弾発射）
- [x] 敵インベーダーの波状出現と移動
- [x] 衝突判定（弾と敵、敵弾と自機）
- [x] 重力波シールド（敵弾を跳ね返す）
- [x] スコア表示とゲームオーバー判定
- [x] Go HTTP サーバーによる静的ファイル配信
- [x] Health endpoint (`/health`)
- [x] Docker コンテナ化
- [x] Helm Chart による Kubernetes デプロイ
- [x] GitHub Actions CI

## 使用方法

### ローカル実行

```bash
# 依存関係をダウンロード
go mod download

# サーバー起動
go run ./server/cmd/server/

# ブラウザでアクセス
open http://localhost:8080
```

### Docker 実行

```bash
# ビルド
docker build -t starhopper .

# 実行
docker run -p 8080:8080 starhopper
```

### Kubernetes へのデプロイ

```bash
# Helm Chart でデプロイ
helm install starhopper ./charts/starhopper

# ポートフォワードでアクセス
kubectl port-forward svc/starhopper 8080:8080
```

## 検証コマンド

```bash
# テスト実行
go test ./...

# リントチェック
golangci-lint run

# 健康状態確認
curl http://localhost:8080/health

# ゲーム画面確認
curl http://localhost:8080/
```

## リポジトリ構成

```
├── client/          # フロントエンド（HTML/CSS/JS）
├── server/          # Go HTTP サーバー
├── charts/          # Helm Chart
├── docs/            # ドキュメント
│   ├── artifact-contract.md
│   └── sprint-1-report.md
├── Dockerfile
├── go.mod
└── README.md
```

## 既知の制限事項

- 敵インベーダーのタイプは現在 1 種類のみ
- シールドのクールダウンは実装済みだが、視覚的なクールダウン表示は未実装
- サウンドエフェクトは未実装
- マルチプレイヤーは未実装

## 次の Sprint で実装予定

- [ ] 複数の敵インベーダータイプ（異なる動き・HP）
- [ ] シールドのクールダウン視覚化
- [ ] サウンドエフェクト
- [ ] ハイスコアの永続化（localStorage）
- [ ] レベルシステム

## ライセンス

MIT License
