# 🚀 Nova Invaders: 次世代インベーダーゲーム

**Nova Invaders** は、従来のインベーダーゲームの枠を超えた、**「軌道予測と重力波シールド」**を核とした新規性のあるアーケードゲームです。プレイヤーは敵の弾道を予測し、重力波を操って防御・反撃するコアループを体験します。

## 🎮 プロダクト概要

- **対象ユーザー**: アーケードゲームファン、レトロゲーム愛好家、シンプルかつ戦略的なアクションを好むプレイヤー
- **コアループ**: 敵の出現 → 弾道予測 → 重力波シールド展開 → 反撃 → スコア加算
- **差別化要素**: 
  - **重力波シールド**: 敵弾を吸収してエネルギーに変換し、強力な特殊攻撃として使用可能
  - **軌道予測インジケーター**: 敵の弾道が可視化され、戦略的な防御が可能
- **Non-Goals**: マルチプレイヤー、複雑なストーリーモード、外部サービス連携

## 🚀 クイックスタート

### ローカル実行

```bash
# 依存関係のインストール
cd server && go mod download

# サーバー起動
go run main.go

# ブラウザでアクセス
open http://localhost:8080
```

### Docker 実行

```bash
# イメージビルド
docker build -t nova-invaders .

# コンテナ起動
docker run -p 8080:8080 nova-invaders
```

## ✅ 検証コマンド

### スモークテスト

```bash
# 健康チェックエンドポイント
curl -f http://localhost:8080/health

# メインページアクセス
curl -f http://localhost:8080/

# ゲームアセット検証
curl -f http://localhost:8080/assets/game.js
curl -f http://localhost:8080/assets/style.css
```

### テスト実行

```bash
# ユニットテスト
cd server && go test ./...

# 構文チェック
cd server && go vet ./...
```

## 📦 デプロイメント

### Kubernetes (Helm)

```bash
# Helm チャートインストール
helm install nova-invaders ./charts/nova-invaders \
  --set image.repository=nova-invaders \
  --set image.tag=latest

# ポートフォワード
kubectl port-forward svc/nova-invaders 8080:80
```

## 📁 リポジトリ構成

```
.
├── server/          # Go HTTP サーバーとバックエンドロジック
│   ├── main.go      # エントリポイント
│   ├── handlers/    # HTTP ハンドラー
│   ├── game/        # ゲームドメインロジック
│   └── go.mod       # Go モジュール定義
├── client/          # フロントエンドアセット
│   ├── index.html   # メインページ
│   ├── assets/      # CSS/JS/画像
│   └── game.js      # ゲームロジック
├── charts/          # Helm チャート
│   └── nova-invaders/
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
├── docs/            # プロダクトドキュメント
│   └── artifact-contract.md
├── Dockerfile       # コンテナビルド定義
├── .github/         # CI/CD ワークフロー
└── README.md        # このファイル
```

## ⚠️ 既知の制限事項

1. **シングルプレイヤーのみ**: マルチプレイヤー機能は未実装
2. **ローカルストレージ**: スコアデータはブラウザのローカルストレージに保存
3. **モバイル最適化未完了**: デスクトップファーストの実装
4. **外部サービス非対応**: 認証、スコアランキング、分析などは未実装

## 📋 次のスプリントへのバックログ

- [ ] マルチプレイヤー対応
- [ ] モバイルレスポンシブデザイン
- [ ] スコアランキングシステム
- [ ] アチーブメントシステム
- [ ] パフォーマンス最適化
- [ ] 詳細なモーションアニメーション

## 🔍 検証ポイント

- [x] メインページが正常にレンダリングされる
- [x] ゲームアセットが正しく読み込まれる
- [x] 健康チェックエンドポイントが正常に動作する
- [x] Docker コンテナが正常に起動する
- [x] Helm チャートが正常にデプロイされる
- [x] ゲームコアループが動作する

---

**Nova Invaders** は、シンプルでありながら戦略的な深みを持つゲーム体験を提供します。重力波シールドと軌道予測という新規性のあるメカニクスで、インベーダーゲームの新しい可能性を提示します。