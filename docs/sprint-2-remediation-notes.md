# Sprint 2 修正ノート (Remediation Notes)

## 概要
Sprint 2 QA 実施中に確認されたデプロイメントおよびパッケージ関連の失敗事項を整理し、Sprint 2 チェックポイント前に実行する修正計画を記載します。

## 発見された失敗事項と原因分析

### 1. Docker ビルド失敗
- **現象**: `docker build` 実行時に依存関係の解決エラーまたはキャッシュミスが発生。
- **原因**: `Dockerfile` 内の `COPY` 順序が最適化されておらず、依存関係インストールステップが頻繁に再実行されている。または、Go モジュールキャッシュが適切にマウントされていない。
- **修正方針**: `COPY go.mod go.sum` を先に配置し、依存関係インストールをキャッシュ層として分離。`go build` 時に環境変数の設定を確認。

### 2. Helm Chart / Kubernetes マニフェスト検証エラー
- **現象**: `helm lint` または `kubectl apply --dry-run` でスキーマ違反またはリソース不足エラー。
- **原因**: Deployment の `resources` セクションに `requests` が未定義、または Service の `selector` と Deployment の `labels` が一致していない。
- **修正方針**: `charts/` 内の `values.yaml` と `deployment.yaml` を照合し、ラベルセレクタを統一。リソース制限（CPU/Memory）の最小値を定義。

### 3. CI パイプライン (GitHub Actions) 失敗
- **現象**: `go test` または `golangci-lint` ステージでタイムアウトまたはパスエラー。
- **原因**: CI 環境でのテスト並列実行設定不足、または lint ルールが新規コードに厳しすぎる。
- **修正方針**: `.github/workflows/` の `go test` に `-parallel` オプションを適用。lint ルールをプロジェクト規模に合わせて調整し、必須チェックのみを有効化。

## 修正計画 (Remediation Steps)
1. `Dockerfile` のレイヤー最適化とキャッシュ戦略の見直し。
2. Helm Chart の `values.yaml` とマニフェストの整合性チェック、リソース定義の追加。
3. CI ワークフローのテスト並列化と lint ルールの調整。
4. 各修正後、ローカルで `docker build`, `helm lint`, `go test ./...` を実行し、成功を確認。

## 検証方法
- ローカル環境で `docker build -t <image> .` が完了すること。
- `helm lint charts/` が警告なしで完了すること。
- `go test -v ./...` が全テストパスすること。
- CI 環境でワークフローが緑色になることを確認。

## 影響範囲とリスク
- 修正はビルドパイプラインとデプロイメントアーティファクトに限定。
- 製品ロジックや UI への影響はない。
- 修正完了後、Sprint 2 チェックポイント前に QA 環境で再検証を実施。

## 備考
- 本ドキュメントは計画段階の产物であり、実装は次ステップで実施。
- 親タスクの指示に従い、不要なアーティファクトや全文コピーは含めない。