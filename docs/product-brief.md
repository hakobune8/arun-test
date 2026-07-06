# Hakobune Invaders — Product Brief

## Concept
**Hakobune Invaders** は、プレイヤーが小舟（はこぶね）を操り、空から襲いかかる敵を撃退するインベーダーゲームです。従来のインベーダーゲームとの差別化として、**舟の揺れ（波のシミュレーション）**と**弾の流され（潮流エフェクト）**を組み合わせた、直感的でありながら戦略的な操作体験を提供します。

## Target User
- シンプルなアーケードゲームを好むカジュアルゲーマー
- 日本語圏のゲーム愛好家
- 5〜10 分のプレイセッションを求めるユーザー

## Core Loop
1. **敵の出現** — 画面上部から敵が波状に出現
2. **プレイヤー操作** — 矢印キーで舟を左右に移動（波の影響で自然な揺れが発生）
3. **攻撃** — スペースキーで弾を発射（弾は潮流の影響でわずかに流される）
4. **スコア計算** — 敵を撃墜するとスコア加算、コンボボーナス付き
5. **ライフ管理** — 敵の弾が舟に命中するとライフ減少、ライフ 0 でゲームオーバー
6. **リトライ** — ゲームオーバー後、スコア表示とともに即座にリトライ可能

## Differentiating Behavior
- **舟の揺れ**: 舟は単に左右に動くだけでなく、波の高さに応じて上下にも揺れる。これにより、敵の弾を回避する際に「タイミング」が重要になる。
- **潮流エフェクト**: プレイヤーの弾は潮流の影響でわずかに流される。敵の位置を読みながら弾を放つ必要がある。
- **ポップなビジュアル**: 明るいカラーパレット、シンプルなドット絵スタイル、アニメーション付きの爆発エフェクト。

## Acceptance Criteria
| # | 要件 | 観測可能な検証方法 |
|---|------|-------------------|
| AC-1 | `/healthz` エンドポイントが `200 OK` を返す | `curl http://localhost:8080/healthz` でステータスコードを確認 |
| AC-2 | `/` ルートでゲーム UI が表示される | ブラウザでアクセスし、ゲームタイトルと操作パネルが表示される |
| AC-3 | 舟は左右キーで移動可能 | キー入力に対して舟が画面内で移動する |
| AC-4 | スペースキーで弾が発射される | キー入力に対して弾が画面内を移動する |
| AC-5 | 敵を撃墜するとスコアが加算される | ゲームプレイ中にスコア表示が更新される |
| AC-6 | ゲームオーバー時にスコアとリトライボタンが表示される | ライフ 0 になるとスコア画面に遷移する |
| AC-7 | Dockerfile が存在し、`docker build` が成功する | `docker build -t hakobune-invaders .` がエラーなく完了する |
| AC-8 | Helm chart が存在し、`helm template` が成功する | `helm template hakobune-invaders charts/hakobune-invaders/` が YAML を出力する |
| AC-9 | GitHub Actions CI が build と test を実行する | PR 作成時に CI がパスする |

## Non-Goals
- マルチプレイヤー対応（Sprint 1 ではシングルプレイのみ）
- 外部サービス連携（スコア保存、認証など）
- モバイル対応（Sprint 1 ではブラウザデスクトップのみ）
- 本格的なサウンドエフェクト（Sprint 1 では視覚のみに集中）

## Qualitative Requirements → Observable Criteria
| 定性要求 | 観測可能な criteria |
|----------|-------------------|
| 新規性 | 舟の揺れと潮流エフェクトが実装され、ゲームプレイに影響を与える |
| 楽しい | コントレスポンスが 200ms 以下、ゲームオーバーからリトライまで 1 秒以内 |
| ポップ | 明るいカラーパレット（背景: #1a1a2e, プレイヤー: #e94560, 敵: #0f3460）、アニメーション付き |
| シンプル | 操作は矢印キー（移動）とスペースキー（発射）のみ、UI に説明文を表示 |
| Production-ready | `/healthz` エンドポイント、Dockerfile、Helm chart、CI pipeline が存在 |

## Sprint 1 Scope
- Go net/http サーバーの実装（`/healthz`, `/` ルート）
- 最小限のゲーム UI（HTML/CSS/JS）
- Dockerfile と Helm chart の基本構成
- GitHub Actions CI の基本 pipeline
- README と検証コマンド

## File Layout (Sprint 1)
```
.
├── docs/
│   ├── product-brief.md          ← このファイル
│   └── artifact-contract.md      ← 後続ステップで作成
├── server/
│   ├── main.go                   ← Go エントリポイント
│   ├── handler/                  ← HTTP ハンドラー
│   └── config/                   ← 設定
├── client/
│   ├── index.html                ← ゲーム UI エントリポイント
│   ├── css/                      ← スタイルシート
│   └── js/                       ← ゲームロジック
├── charts/
│   └── hakobune-invaders/        ← Helm chart
├── Dockerfile
├── go.mod
├── README.md
└── .github/workflows/            ← CI pipeline
```
