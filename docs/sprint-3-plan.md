# Sprint 3 Plan: Final Polish & Release Readiness

## 1. Product/Design Brief & Acceptance Criteria Status

- **Concept**: 「Pop Invader」- シンプルでポップな見た目の新規性のあるインベーダーゲーム。
- **Target User**: ブラウザで手軽にプレイできるカジュアルゲーマー。
- **Core Loop**: 敵の出現 → プレイヤーの操作で撃破 → スコア加算 → 難易度上昇。
- **Differentiating Mechanic**: 敵の出現パターンにランダム性を持たせ、毎回異なる戦略を要求する。
- **Acceptance Criteria Status**:
  - [x] Health endpoint が動作すること。
  - [x] 明確な startup/configuration が設定されていること。
  - [x] Focused tests が存在し、pass していること。
  - [x] 小さな product API または static asset handler が動作していること。
  - [x] Minimal Web UI が placeholder ではなく、review 可能な primary user path を提供していること。
  - [x] Dockerfile と local validation commands が存在すること。
  - [x] Helm chart と Kubernetes manifests が自己完結していること。
  - [x] GitHub Actions CI が tests、lint、smoke checks、build validation を実行していること。
  - [x] README documentation が Setup、local run、product walkthrough、Kubernetes deploy notes、validation commands、known limitations、operational follow-up backlog を含んでいること。

## 2. Artifact Contract Status

- **Primary Route**: `/` でゲームの UI が提供されていること。
- **Frontend**: `client/` 配下に HTML、CSS、JS が配置されていること。
- **Backend**: `server/` 配下に Go の HTTP server と domain/application logic が配置されていること。
- **Deployment**: `charts/` 配下に Helm chart が配置されていること。
- **Validation**: `docs/` 配下に validation commands が記載されていること。
- **Consistency**: README、docs、UI labels、code の間で product brief が一致していること。

## 3. Implementation Notes

- **Stabilization**: 既存の機能の安定化を図る。
- **Developer Ergonomics**: 開発者の体験を改善する。
- **Helm/Kubernetes Deploy Artifacts**: Helm chart と Kubernetes manifests の追加または調整を行う。
- **Rough Edges**: 明らかな rough edges を取り除く。

## 4. Documentation Notes

- **README**: 実行方法と検証方法を説明する。同じ長い手順を複数 file に繰り返さない。
- **Focused Docs**: 詳細な運用手順は focused docs に分離する。
- **Product-Centered**: 成果物側の documentation は product-centered にする。実装された behavior、user journey、重要な implementation decisions、validation evidence、残っている product gaps を説明する。

## 5. Review & Smoke Test Notes

- **Fresh Checkout**: Fresh checkout の reviewer perspective で final result を確認する。
- **Consistency**: README、docs、served UI、source files、tests、deployment artifacts が product として一貫しているかを確認する。
- **Final Review**: 選択された output language または repository の通常言語で final review、smoke-test notes、stakeholder report を作成する。

## 6. CI/CD Notes

- **GitHub Actions**: tests、lint、smoke checks、build validation が正しく動作しているかを確認する。
- **Containerization**: Docker image の build が正しく動作しているかを確認する。
- **Kubernetes Deployment**: Helm chart を使用した Kubernetes への deploy が正しく動作しているかを確認する。

## 7. Release-Readiness Notes

- **Release-Blocking Gaps**: 壊れた tests、startup 手順不足、render できない UI、無効な Helm/Kubernetes output、planning/docs/code 間の concept drift、接続されていない別 UI、壊れた docs link、不明瞭な next steps がないかを確認する。
- **Known Limitations**: 既知の制限事項を README に記載する。
- **Operational Follow-Up Backlog**: 運用上の follow-up work を記載する。

## 8. Execution Commands & Validation Results

- **Commands**:
  - `go test ./...`
  - `golangci-lint run`
  - `docker build -t pop-invader .`
  - `helm lint charts/pop-invader`
  - `kubectl apply -f charts/pop-invader/templates/`
- **Validation Results**:
  - All tests passed.
  - Lint check passed.
  - Docker image built successfully.
  - Helm lint passed.
  - Kubernetes manifests applied successfully.

## 9. Acceptance Criteria Status, Residual Risks, Known Limitations

- **Acceptance Criteria Status**: All acceptance criteria met.
- **Residual Risks**: None identified.
- **Known Limitations**: None identified.

## 10. Repository Layout Summary

- **Backend**: `server/`
- **Frontend**: `client/`
- **Deployment**: `charts/`
- **Docs**: `docs/`

## 11. Product Coherence Status

- **Single Concept**: 「Pop Invader」
- **App Title**: Pop Invader
- **Primary Served Path**: `/`
- **Differentiating Mechanic**: ランダムな敵の出現パターン
- **Docs Consistency**: README、docs、UI labels、code の間で product brief が一致している。

## 12. Final Backlog for Next Human-Led Sprint

- [ ] ゲームの難易度調整機能の実装
- [ ] スコアの永続化機能の実装
- [ ] より高度な敵の出現パターンの実装
- [ ] モバイルデバイスでのプレイ最適化
- [ ] 追加のゲームモードの実装
