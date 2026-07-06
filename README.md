# Novel Invader Game

新規性のあるインベーダーゲームです。ポップでシンプルなUIと、従来のインベーダーゲームとは異なる「弾の軌道予測」メカニクスを搭載しています。

## 概要
- **対象ユーザー**: アーケードゲームファン、カジュアルゲーマー
- **コアループ**: 敵の出現 → 弾の軌道予測 → 回避/撃破 → スコア加算
- **差別化要素**: 弾の軌道を可視化し、プレイヤーが戦略的に回避・撃破できるUI
- **非目標**: 複雑なストーリーモード、マルチプレイヤー、外部サービス連携

## Helm によるデプロイ

このアプリケーションは、ARUN と同じ Kubernetes 環境にデプロイできるように設計されています。Helm チャートを使用して簡単に展開できます。

### 前提条件
- `helm` v3 以上
- `kubectl` がクラスタに接続されていること

### インストール
```bash
helm install novel-invader ./charts/novel-invader \
  --set image.repository=<your-registry>/novel-invader \
  --set image.tag=<version>
```

### 値のカスタマイズ
`charts/novel-invader/values.yaml` を編集するか、`--set` フラグで以下を設定できます：
- `image.repository`, `image.tag`: コンテナイメージ
- `service.type`: Service タイプ (`ClusterIP`, `NodePort`, `LoadBalancer`)
- `resources.limits`, `resources.requests`: リソース制限
- `livenessProbe`, `readinessProbe`: ヘルスチェック設定

### 検証
```bash
kubectl get pods -l app.kubernetes.io/name=novel-invader
kubectl get svc -l app.kubernetes.io/name=novel-invader
kubectl logs -l app.kubernetes.io/name=novel-invader --tail=20
```

詳細なデプロイ手順や運用ガイドは [docs/deployment.md](docs/deployment.md) を参照してください。