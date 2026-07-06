# Gravity Flip Invaders

**新規性のあるインベーダーゲーム** - 重力反転メカニクスを搭載したレトロシューティングゲーム。

## 🚀 プロダクト概要

このプロジェクトは、従来の上下左右移動型インベーダーゲームとは異なる、**「重力反転（Gravity Flip）」**という差別化されたメカニクスを持つゲームです。プレイヤーは画面下部の宇宙船を操作し、敵の弾を回避しつつ攻撃します。最大の特徴は、`G`キー（またはボタン）で重力を反転させ、画面の天井と床を行き来できる点です。これにより、単純な横移動だけでなく、立体的な回避行動と攻撃角度の制御が可能になります。

### 対象ユーザー
- レトロゲームを愛するプレイヤー
- シンプルだが奥深いメカニクスを求めるカジュアルゲーマー

### コアループ
1. **移動**: 矢印キーで左右に移動。
2. **重力反転**: `G`キーで重力を反転（天井移動 ↔ 床移動）。
3. **攻撃**: スペースキーで弾を発射。重力方向に弾が飛ぶ。
4. **クリア**: 敵を全滅させる。

## ✅ Sprint 1 達成状況 (Acceptance Criteria)

| 項目 | 状態 | 備考 |
| :--- | :--- | :--- |
| **新規性 (Gravity Flip)** | ✅ 実装済み | 重力反転による立体的な移動と弾の軌道制御 |
| **Health Endpoint** | ✅ 実装済み | `/health` でサーバー状態を確認可能 |
| **Static Asset Serving** | ✅ 実装済み | `client/` 配下の UI を Go サーバー経由で提供 |
| **Dockerfile** | ✅ 実装済み | 多段ビルドによる軽量イメージ |
| **Helm Chart** | ✅ 実装済み | Kubernetes 環境へのデプロイ対応 |
| **CI/CD** | ✅ 実装済み | GitHub Actions による lint/test/build |

## 🛠️ セットアップと実行

### ローカル実行 (Go)

```bash
# 依存関係の取得
go mod tidy

# サーバー起動 (frontend assets も含めてローカルで動作します)
go run ./server/...
```

ブラウザで `http://localhost:8080` にアクセスしてゲームを開始できます。

### Docker 実行

```bash
# イメージのビルド
docker build -t gravity-flip-invaders .

# コンテナの実行
docker run -p 8080:8080 gravity-flip-invaders
```

### Kubernetes (Helm) へのデプロイ

```bash
# Helm チャートのインストール
helm install gravity-flip ./charts/gravity-flip --set image.repository=your-registry/gravity-flip-invaders --set image.tag=latest
```

## 📂 リポジトリ構成

```
.
├── server/          # Go HTTP サーバーとバックエンドロジック
│   ├── main.go      # エントリポイント
│   └── ...
├── client/          # フロントエンド (HTML/CSS/JS)
│   ├── index.html   # ゲーム画面
│   └── ...
├── charts/          # Helm Chart (Kubernetes 用)
├── docs/            # プロダクト仕様、契約、レポート
│   ├── artifact-contract.md
│   └── sprint-report.md
├── Dockerfile       # コンテナビルド定義
└── README.md        # このファイル
```

## 📝 既知の制限事項 (Known Limitations)

- **サウンド**: 現時点では効果音・BGMは実装していません。
- **スコア保存**: リロードするとスコアはリセットされます（永続化は未実装）。
- **モバイル対応**: キーボード操作が前提のため、モバイルブラウザでのプレイは想定していません。

## 🔜 次の Sprint へ

- サウンドエフェクトの追加
- ハイスコアのローカルストレージへの保存
- モバイルタッチ操作のサポート
