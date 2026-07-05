# Space Drift: Gravity Invader

新規性のあるインベーダーゲーム。重力場を操作して敵を倒す、シンプルながら戦略的なアーケード体験。

## 概要
本プロジェクトは、従来のインベーダーゲームの枠組みに「重力場シフト」メカニズムを追加した、Go 製軽量 Web ゲームです。プレイヤーは自機を操作するだけでなく、画面内の重力方向を切り替えることで敵の軌道を変え、効率的に撃破します。

## クイックスタート
```bash
# ローカル実行
go run ./cmd/server/main.go

# Docker 実行
docker build -t space-drift .
docker run -p 8080:8080 space-drift
```

## ユーザー体験
- **コアループ**: 自機移動 → 重力場シフト → 敵軌道変化 → 撃破 → スコア加算
- **差別化要素**: 物理演算に基づく軌道制御。単純な射撃だけでなく、環境利用が勝利の鍵。
- **UI**: シンプルな Canvas ベースのレンダリング。レスポンシブ対応。

## 検証エビデンス
- [Sprint 1 Report](docs/sprint-1-report.md) に QA 結果、リメデーション、バックログを記載。
- CI パイプライン: `make test` および `make lint` が main ブランチで自動実行。
- K8s デプロイ: Helm chart 経由で Service/Deployment/Probes 付きで展開可能。

## 既知の制限事項
- 単一プレイヤーモードのみ実装済み。マルチプレイは次回 Sprint で検討。
- 外部アセットは最小限に抑え、ローディング時間を排除。
- 難易度調整は固定値。動的調整は次回 Sprint で実装予定。

## ドキュメント
- [Sprint 1 レポート](docs/sprint-1-report.md)
- [Kubernetes デプロイ手順](docs/k8s-deploy.md)