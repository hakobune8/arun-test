# Gravity Invaders — Product Brief

## Concept

**Gravity Invaders** is a browser-based space invader game where the player can **flip gravity** to move between the top and bottom of the screen, dodging enemy fire and positioning for optimal shots.

## Target User

Casual browser gamers who enjoy arcade-style games with a quick learning curve and a single novel mechanic that creates fresh tactical decisions.

## Core Loop

1. Player ship moves left/right and can **flip gravity** (spacebar) to switch between top and bottom lanes.
2. Enemy waves descend from the center, firing in patterns.
3. Player shoots enemies while avoiding collisions and incoming fire.
4. Clearing a wave triggers the next wave with increased difficulty.
5. Game ends when the player loses all lives.

## Differentiating Mechanic

**Gravity Flip** — Pressing spacebar instantly reverses the player's vertical position (top ↔ bottom of the screen). This is not a gradual movement; it is an instant toggle that creates a binary lane system. The mechanic must be visible in the source code as a distinct function (`FlipGravity`) and in the UI as an instant visual transition.

## Qualitative Requirements → Observable Criteria

| Qualitative Term | Observable Criterion |
|---|---|
| 新規性 (Novel) | Gravity flip mechanic implemented as a distinct toggle, not present in standard invader games |
| 楽しい (Fun) | Smooth 60fps gameplay, responsive controls (<100ms input lag), satisfying visual feedback on hits |
| ポップ (Playful) | Bright color palette, simple particle effects on enemy destruction, cheerful sound cues |
| シンプル (Simple) | Single screen, no menus beyond start/game over, one primary control (spacebar for gravity flip) |
| Production-ready | Runs locally via `go run`, builds via Docker, deploys via Helm chart, passes CI checks |

## Non-Goals

- Multiplayer or online leaderboards
- Complex power-ups or item systems
- Persistent save states or account systems
- Mobile touch controls (keyboard only for Sprint 1)
- Complex audio (placeholder sounds or silent for Sprint 1)

## Sprint 1 Acceptance Criteria

- [ ] Game renders in browser at `http://localhost:8080/` with playable invader game
- [ ] Gravity flip mechanic works: spacebar toggles player between top and bottom lanes
- [ ] Enemies spawn in waves and descend toward the player
- [ ] Player can shoot and destroy enemies
- [ ] Collision detection works (player loses life on enemy contact)
- [ ] Game over screen displays score and restart option
- [ ] Go server starts cleanly with health endpoint at `/health`
- [ ] Docker build succeeds and container runs
- [ ] Helm chart deploys to Kubernetes (Service, Deployment, probes)
- [ ] GitHub Actions CI passes (build, test, lint)
- [ ] README documents how to run, validate, and deploy

## Product Title

**Gravity Invaders** (used consistently in UI, docs, and code comments)

## Primary Served Path

`GET /` → serves the game UI (HTML + embedded CSS/JS or static assets from `client/`)

## Repository Layout

```
├── server/          # Go HTTP server and backend logic
├── client/          # Browser assets (HTML, CSS, JS)
├── docs/            # Product brief, artifact contract, validation notes
├── charts/          # Helm chart for Kubernetes deployment
├── Dockerfile       # Container build
├── go.mod           # Go module definition
├── README.md        # Entry point documentation
└── .github/         # CI workflows
```
