# Sprint 2 レポート

## 概要
Sprint 2 では、Sprint 1 で確立した最小限の Go HTTP サーバーとフロントエンドスライスに、本格的なゲームロジック、テスト、CI/CD パイプライン、および Kubernetes 展開アーティファクトを追加しました。

## 展開アーティファクト (Deployment Artifacts)
- **Dockerfile**: 多段ビルドを採用し、最終イメージを最小限に抑えました。`/health` エンドポイントを含むサーバーをベースにしています。
- **Helm Chart**: `charts/` ディレクトリに配置。`values.yaml` でリソース制限とレプリカ数を設定可能にしました。
- **Kubernetes Manifests**: `k8s/` ディレクトリに配置。Deployment（レプリカ、プローブ付き）、Service（ClusterIP）を含みます。

## QA 証拠 (QA Evidence)
- **Unit Tests**: `go test ./...` が全パスしました。ゲームロジック（弾の移動、衝突判定）の単体テストを追加しました。
- **Smoke Tests**: `docker build` および `docker run` によるローカル検証が成功しました。
- **Helm Validation**: `helm template .` により、マニフェストの生成が正常に動作することを確認しました。

## 実施した修正 (Remediation Performed)
- **ポート競合の修正**: サーバーが環境変数 `PORT` でポートを動的に受け付けるように修正しました。
- **Helm チャートのセレクター不整合**: Deployment の `selector.matchLabels` と `template.metadata.labels` の不一致を修正しました。
- **ヘルスチェックエンドポイント**: `/health` エンドポイントを追加し、Kubernetes の liveness/readiness プローブに対応しました。

## 残りの CI/リリース バックログ (Remaining CI/Release Backlog)
- **統合テスト**: フロントエンドとバックエンドの連携テストを追加する。
- **CI パイプラインの強化**: `helm lint` や `golangci-lint` などの静的解析ツールを CI に追加する。
- **パフォーマンスチューニング**: ゲームループのフレームレート安定化とメモリ使用量の最適化。
- **Ingress 設定**: 外部からのアクセスを可能にする Ingress リソースの追加（Sprint 3 で検討）。

## 製品の整合性 (Product Coherence)
- **コンセプト**: 「新規性のあるインベーダーゲーム」は、従来のインベーダーゲームに「弾の軌道予測」 mechanic を追加したもので統一されています。
- **UI/ラベル**: ゲームタイトル、UI ラベル、README すべてで同じコンセプトを参照しています。
- **主要なユーザーパス**: ブラウザでゲームにアクセスし、プレイできることが確認されています。

## 次のステップ
- Sprint 3 では、残りのバックログを消化し、最終的な安定化とドキュメントの完成を目指します。
