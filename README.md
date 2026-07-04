# Gravity Invaders

重力を操作して敵を倒す、ミニマルなインベーダーゲーム。

## 概要

**Gravity Invaders** は、従来のインベーダーゲームの枠組みに「重力反転」という差別化されたメカニクスを加えた、新規性のあるアーケードゲームです。

- **対象ユーザー**: クラシックゲームを好むが、新しいインタラクションを求めるプレイヤー。
- **コアループ**: 敵の弾を回避しつつ、重力を反転させて自機の位置を移動し、敵を撃破する。
- **差別化ポイント**: 「重力反転」による直感的な操作と、シンプルながら奥深い回避アクション。

詳細な製品要件は [docs/product-brief.md](docs/product-brief.md) を参照してください。

## セットアップ

### ローカル実行

Go 1.21 以上が必要です。

```bash
# ビルドと実行
go run ./cmd/server/

# ブラウザで確認
open http://localhost:8080
```

### Docker

```bash
docker build -t gravity-invaders .
docker run -p 8080:8080 gravity-invaders
```

### Kubernetes (Helm)

```bash
helm install gravity-invaders ./charts/gravity-invaders
```

## 検証

### スモークテスト

```bash
# テスト実行
go test ./...

# ヘルスチェック
curl -f http://localhost:8080/health
```

### 製品検証

1. ブラウザで `http://localhost:8080` にアクセスします。
2. ゲーム画面が表示され、タイトルと「Start」ボタンが表示されることを確認します。
3. ゲームを起動し、重力反転（スペースキーまたはクリック）で操作できることを確認します。

## リポジトリ構成

```
.
├── cmd/server/          # アプリケーションのエントリポイント
├── internal/            # 内部パッケージ（ロジック、ゲームエンジン）
├── frontend/            # 静的アセット（HTML, CSS, JS）
├── charts/              # Helm チャート
├── docs/                # 製品ドキュメント
├── Dockerfile           # コンテナビルド定義
└── README.md            # このファイル
```

## 既知の制限

- サウンドエフェクトは実装されていません。
- スコアの永続化（データベース連携）は行われていません。
- モバイルブラウザでのタッチ操作は最適化されていません。

## 次のステップ

- サウンドエフェクトの実装
- スコアの永続化
- モバイル対応の最適化
- CI/CD パイプラインの強化
