# Neon Invaders — Product Brief

## Concept

**Neon Invaders** は、レトロなインベーダーゲームの美学を、現代的なネオン美学とシンプルな操作で再解釈したブラウザゲームです。

- **対象ユーザー**: レトロゲームを愛するカジュアルゲーマー、ブラウザで手軽にプレイしたいユーザー
- **コアループ**: 敵を撃ち落とす → スコア獲得 → 難易度上昇 → より高いスコアを目指す
- **差別化 mechanic**: 敵が「光の輪」を形成し、プレイヤーは輪の隙間を縫って攻撃する。敵の配置パターンが各レベルで変化し、予測不能な戦略的プレイを要求する
- **非目標**: 複雑なストーリー、マルチプレイヤー、モバイル最適化、外部サービス連携

## Acceptance Criteria

| 定性要求 | 観測可能な基準 |
|---------|--------------|
| 新規性 | 敵が「光の輪」パターンで出現し、プレイヤーが隙間を縫う戦略が必要 |
| 楽しい | 1プレイで30秒〜2分のゲームセッションが可能、スコア表示で達成感 |
| ポップ | ネオンカラー（シアン、マゼンタ、イエロー）を基調としたビジュアル |
| シンプル | キーボード操作のみ（←→で移動、スペースで射撃）、30秒以内に理解可能 |
| Production-ready | 外部依存なし、ブラウザで直接動作、エラー時に graceful degradation |

## Core Loop

1. プレイヤーが画面下部で左右に移動
2. 画面上部から「光の輪」パターンで敵が出現
3. スペースキーで弾を発射、敵を撃ち落とす
4. 全敵撃破で次のレベル（敵の速度・パターンが変化）
5. プレイヤーが敵の弾に当たるとゲームオーバー

## Implementation Notes

- Canvas-based rendering for smooth animation
- No external dependencies — pure HTML/CSS/JS frontend
- Go backend serves static assets and provides health check
- Single-page application, no routing complexity
