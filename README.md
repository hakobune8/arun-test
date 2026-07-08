# Gravity Invaders

**重力で操る、新しいインベーダーゲーム**

## 1. Product Overview

### Concept
従来の左右移動型インベーダーゲームとは異なり、**「重力（Gravity）」**を操作して自機を浮遊・落下させながら敵弾を回避し、攻撃するゲームです。

- **Target User**: レトロゲームを愛するプレイヤー、新しい操作感を探求するゲーマー。
- **Core Loop**: 重力を制御して敵弾を回避 → 敵を撃破 → スコア獲得 → 次の波へ。
- **Differentiating Mechanic**: マウス/タッチの上下で重力方向を制御。直感的でありながら、弾幕回避の奥深さを持つ。
- **Non-Goals**: 複数プレイヤー、オンラインランキング、複雑なストーリーモード。

### Acceptance Criteria (Sprint 1)
- [x] ブラウザで動作する単一画面のゲームUIを提供する。
- [x] 重力操作による自機の移動と、弾の発射が動作する。
- [x] 敵の出現と弾の発射、衝突判定が正常に動作する。
- [x] ゲームオーバーとリスタートのフローが確立されている。

## 2. Getting Started

### Local Development
Go 1.21+ が必要です。

```bash
# ビルドと実行
go run ./server/cmd/main.go

# ブラウザで確認
open http://localhost:8080
```

### Docker
```bash
docker build -t gravity-invaders .
docker run -p 8080:8080 gravity-invaders
```

### Kubernetes (Helm)
```bash
helm install gravity-invaders ./charts/gravity-invaders
```

## 3. Validation & QA

### Smoke Test
```bash
# Health check
curl -f http://localhost:8080/health

# Game UI access
curl -f http://localhost:8080/
```

### Unit Tests
```bash
go test ./...
```

## 4. Known Limitations & Backlog

### Known Limitations
- **Graphics**: プログラミングによる簡易描画（Canvas API）を使用。アセット読み込みは行っていない。
- **Audio**: 現時点では効果音なし。
- **Persistence**: スコアの保存機能は未実装。

### Operational Backlog (Next Sprint)
- [ ] 効果音の追加（Web Audio API）。
- [ ] スコアのローカルストレージへの保存。
- [ ] 敵のバリエーション増加（ボス戦の導入）。
- [ ] モバイル端末でのタッチ操作の最適化。

## 5. Repository Layout

```
.
├── server/          # Go backend & game logic
├── client/          # Frontend assets (HTML/CSS/JS)
├── charts/          # Helm chart for K8s deployment
├── docs/            # Product & validation docs
└── README.md        # This file
```
