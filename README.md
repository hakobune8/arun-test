# Nova Invader

## 概要
**Nova Invader** は、Go とバニラ JavaScript で構築された、新規性のあるインベーダーゲームです。
従来のインベーダーゲームの枠組みを維持しつつ、「重力反転」メカニクスを実装し、プレイヤーに新たな戦略的深みを提供します。

## Sprint 1 成果物概要
Sprint 1 では、以下の最小限の機能実装とインフラストラクチャを完了しました。

- **コアゲームプレイ**: 基本的な移動、射撃、敵の出現ロジック。
- **差別化メカニクス**: 「重力反転」ボタンによる画面上下の反転操作。
- **UI/UX**: レスポンシブなブラウザゲーム画面。
- **インフラ**: Go HTTP サーバー、Dockerfile、Helm Chart、GitHub Actions CI。

## ユーザー体験
1. ブラウザでゲームにアクセス。
2. 矢印キーで移動、スペースキーで射撃。
3. 「重力反転」ボタンを押して敵の軌道と自機の重力方向を反転させ、回避や攻撃のチャンスを作る。

## セットアップと実行

### ローカル実行
```bash
# サーバー起動
go run ./server/cmd/main.go

# ブラウザで http://localhost:8080 にアクセス
```

### Docker 実行
```bash
docker build -t nova-invader .
docker run -p 8080:8080 nova-invader
```

## 検証コマンド

### ヘルスチェック
```bash
curl http://localhost:8080/health
```

### テスト実行
```bash
go test ./...
```

## リポジトリ構成
- `server/`: Go アプリケーション本体
- `client/`: フロントエンドアセット (HTML/CSS/JS)
- `charts/`: Helm チャート
- `docs/`: 設計書および契約定義

## 既知の制限事項
- 音声エフェクトは実装していません。
- 高スコアランキング機能は未実装です。