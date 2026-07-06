# Sprint 2 調整計画: 修正ノート

## 概要
Sprint 2 QA 証拠に基づき、デプロイメントおよびパッケージングに関連する失敗を特定し、Sprint 2 チェックポイント前に修正方針を整理しました。本ドキュメントは計画段階のものであり、コード実装は含まれません。

## 特定された失敗と修正方針

### 1. Docker ビルド失敗
- **原因**: 多段階ビルドのキャッシュミス、ベースイメージの非互換性、または `COPY` パスの不一致。
- **修正方針**: 
  - `Dockerfile` の `FROM` ステージを明確化し、依存関係のインストールとコードコピーの順序を最適化。
  - `.dockerignore` を見直し、不要なファイルのビルド対象からの除外を確認。
  - ローカルで `docker build` を実行し、キャッシュヒット率とビルド時間を検証。

### 2. Helm チャート検証エラー
- **原因**: `values.yaml` とチャートスキーマの不一致、または必須フィールドの欠落。
- **修正方針**:
  - `helm lint` および `helm template` を実行し、マニフェストの妥当性を確認。
  - `values.schema.json` を更新し、必須フィールドと型制約を明確化。
  - リソース制限 (`resources.requests/limits`) のデフォルト値をチャートに追加。

### 3. Kubernetes マニフェストの問題
- **原因**: Deployment のセレクターとラベルの不一致、またはプローブ設定の欠落。
- **修正方針**:
  - `Deployment` の `selector.matchLabels` と `template.metadata.labels` が完全に一致していることを確認。
  - `livenessProbe` と `readinessProbe` を追加し、ヘルスエンドポイントへのパスを正確に設定。
  - ポート名とコンテナポートの定義を `artifact-contract.md` と照合。

### 4. CI パイプラインのギャップ
- **原因**: リント、テスト、ビルド検証のステップが不足している、または失敗時に停止しない。
- **修正方針**:
  - `.github/workflows/ci.yml` に `golangci-lint`、`go test`、`docker build` のステップを追加。
  - 各ステップの `continue-on-error: false` を明示し、失敗時にパイプラインを停止。
  - キャッシュ設定を最適化し、ビルド時間を短縮。

## 影響評価
- 上記の修正は、ローカル開発環境と Kubernetes 展開の両方に影響します。
- 修正完了後、`docs/artifact-contract.md` のバリデーションコマンドを更新し、QA チックリストと整合させます。

## 次のステップ
1. 修正方針に基づき、`Dockerfile`、`charts/`、`k8s/`、`.github/workflows/` のファイルを更新。
2. ローカルで `docker build`、`helm lint`、`kubectl apply --dry-run=client` を実行し、失敗を解消。
3. CI パイプラインを再実行し、すべてのチェックがパスすることを確認。
4. 修正内容を `README.md` および `docs/` に反映し、Sprint 2 チェックポイントに提出。

---
*本ドキュメントは計画段階のものであり、実装は含まれません。*