# Sprint 1 調整計画および是正ノート

## 概要
Sprint 1 の実装と検証プロセスにおいて、以下のギャップと是正事項を特定しました。これらは Sprint 2 の実装およびドキュメント更新の主要な対象となります。

## 1. プロダクト概念と一貫性 (Product Coherence)
- **発見された問題**: 複数のプロダクト概念の混在、またはタイトルと実装の不一致。
- **是正策**:
  - 単一の「Source of Truth」となるプロダクト概念を明確化し、README、UI、コードのすべてで統一する。
  - 「新規性」や「楽しさ」のような定性的な要件を、検証可能な observable criteria に変換する。

## 2. 成果物契約 (Artifact Contract)
- **発見された問題**: `docs/artifact-contract.md` と実際のファイル構成、ルーティング、アセットパスの不一致。
- **是正策**:
  - Primary route、Frontend/Backend のファイル配置、Go module path、Validation commands が実際のリポジトリ状態と一致するように契約文書を更新する。
  - 存在しないファイルやパスへの参照を削除する。

## 3. バックエンド/フロントエンド実装
- **発見された問題**: フロントエンドのプレースホルダー化、または `client/` と `server/` のディレクトリ分離の不徹底。
- **是正策**:
  - プレースホルダーの UI を排除し、レビュー可能な primary user path を提供する lightweight frontend slice を実装する。
  - ブラウザアセットは `client/`、Go サーバーロジックは `server/` に分離し、flat な構成を避ける。
  - Main application path から直接ユーザー体験が提供されることを確認する。

## 4. デプロイメントと CI/CD
- **発見された問題**: Helm chart の断片化、または無効な Kubernetes manifests。
- **是正策**:
  - Helm chart を `charts/<name>/` 配下に自己完結させる（`Chart.yaml`、`values.yaml`、`templates/` を正しく配置）。
  - Service、Deployment、selectors、labels、probes、resource defaults を含む完全な chart を作成する。
  - CI に lint、smoke checks、build validation を追加する。

## 5. ドキュメント
- **発見された問題**: README の冗長化、または focused docs へのリンク不足。
- **是正策**:
  - README は短い入口とし、詳細な手順は focused docs に分離する。
  - 実装された behavior、user journey、validation evidence に焦点を当て、generic process narration を排除する。

## 次のステップ (Sprint 2 Backlog)
1. `docs/artifact-contract.md` の更新と検証。
2. フロントエンドの primary user path 実装。
3. Helm chart の完全化とローカル検証。
4. CI パイプラインの強化。