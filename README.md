# 星屑の防衛線 (Stardust Defense) - 新規インベーダーゲーム

## プロダクト概要
従来のインベーダーゲームとは異なる、**「軌道計算と重力スイング」**を核としたアクションゲームです。プレイヤーは宇宙ステーションを防御し、敵の軌道を読みながら自機の位置を最適化して迎撃します。

## ユーザー体験とコアループ
1. **観測**: 敵の出現パターンと軌道を読み取る。
2. **配置**: 自機を重力の影響を受ける軌道上に配置する。
3. **迎撃**: 敵が近づく瞬間に弾を発射し、軌道共鳴で破壊する。
4. **評価**: 被弾数と消費エネルギーでスコアを算出。

## 差別化される挙動
- **物理演算ベースの弾道**: 単純な直進弾ではなく、重力ポットの影響で曲がる弾道を採用。
- **エネルギー管理**: 無限弾数ではなく、エネルギー回復を待機中に管理する必要がある。

## 受け入れ基準 (Acceptance Criteria)
- [x] ブラウザでゲーム画面がレンダリングされる。
- [x] 自機がマウス/タッチ操作で移動可能。
- [x] 敵が出現し、弾が衝突判定を持つ。
- [x] ゲームオーバー条件（ステーションHP 0）で終了画面へ遷移。
- [x] Docker コンテナ内で起動可能。

## 実行方法

### ローカル実行
```bash
cd server
go run main.go
```
ブラウザで `http://localhost:8080` にアクセス。

### Docker
```bash
docker build -t stardust-defense .
docker run -p 8080:8080 stardust-defense
```

### Kubernetes (Helm)
```bash
helm install stardust-defense ./charts/stardust-defense
```

## リポジトリ構成
- `server/`: Go 製バックエンド (HTTP サーバー、ゲームロジック)
- `client/`: フロントエンド (HTML/CSS/JS)
- `docs/`: 設計書、QA 報告、デプロイメントガイド
- `charts/`: Helm チャート
- `k8s/`: 手動 Kubernetes マニフェスト

## ドキュメント
- [Artifact Contract](docs/artifact-contract.md): 成果物の接続仕様
- [QA Report](docs/qa-report.md): Sprint 1 検証結果
- [Deployment Guide](docs/deployment.md): 運用手順

---
*このプロジェクトは新規性のあるインベーダーゲームの実装を目的としています。*
