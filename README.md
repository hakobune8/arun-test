# 星屑インベーダー (Stardust Invaders)

ポップでシンプル、かつ新規性のあるインベーダーゲーム。プレイヤーは「星屑」を弾として使い、敵を倒しながら宇宙を探索します。

## 何を作ったか

- **Go net/http サーバー**: `/healthz` ヘルスチェック、`/` でフロントエンド提供、`/api/game` ゲーム状態エンドポイント
- **ミニマル Web UI**: Canvas ベースのゲーム画面（CSS/JS 含む）
- **差別化 mechanic**: 「星屑」弾システム — 敵を倒すと星屑が散り、収集するとパワーアップ
- **Docker / Kubernetes 対応**: Dockerfile、Helm chart、K8s manifests 付き
- **CI**: GitHub Actions でテスト・lint・ビルド検証

## 主要なユーザー体験

1. ブラウザで `/` にアクセス → ゲーム画面が即座に表示
2. 矢印キーで移動、スペースで「星屑」発射
3. 敵を倒すと星屑が散る → 収集でスコアUP＆パワーアップ
4. `/healthz` でヘルスチェック可能

## 受け入れ基準 (Acceptance Criteria)

| 定性的意図 | 観測可能な基準 |
|---|---|
| 新規性 | 星屑収集 mechanic が実装され、スコアに反映される |
| 楽しい | ゲームループが60fpsで動作、レスポンス遅延 < 100ms |
| ポップ | CSS でカラフルな配色、アニメーション付き |
| シンプル | 操作は矢印キー＋スペースのみ、UI に不要な要素なし |
| production-ready | `/healthz` が 200 OK、Docker build 成功、Helm install 成功 |

## リポジトリ構成

```
├── client/          # フロントエンド (HTML/CSS/JS)
├── server/          # Go HTTP サーバー & バックエンドロジック
├── charts/          # Helm chart
├── k8s/             # Kubernetes manifests
├── docs/            # 製品仕様・契約書
│   ├── product-brief.md
│   └── artifact-contract.md
├── .github/workflows/  # CI
├── Dockerfile
└── README.md
```

## ローカル実行

```bash
# ビルド & 実行
cd server && go build -o game . && ./game

# または
make run

# ブラウザで確認
open http://localhost:8080
```

## 検証コマンド

```bash
# テスト
go test ./...

# lint
golangci-lint run

# ヘルスチェック
curl -f http://localhost:8080/healthz

# ゲームエンドポイント
curl http://localhost:8080/api/game
```

## Kubernetes デプロイ

```bash
# Helm でデプロイ
helm install stardust-invaders ./charts/stardust-invaders

# ポートフォワード
ekubectl port-forward svc/stardust-invaders 8080:80
```

## 既知の制限

- 単一プレイヤーのみ（マルチプレイ未実装）
- 外部サービス依存なし（オフライン動作可能）
- 音声エフェクト未実装

## 今後の課題

- [ ] 敵の多様なパターン追加
- [ ] ハイスコア保存（ローカルストレージ）
- [ ] モバイルタッチ操作対応
- [ ] 音声エフェクト追加

---

詳細は [`docs/product-brief.md`](docs/product-brief.md) と [`docs/artifact-contract.md`](docs/artifact-contract.md) を参照してください。
