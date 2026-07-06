# Rhythm Invaders — Product Brief

## Concept

**Rhythm Invaders** is a space invaders-style arcade game where gameplay is synchronized to music beats. Enemies appear, move, and can only be effectively destroyed when the player shoots on the beat. This creates a unique fusion of rhythm game precision and classic arcade action.

## Target User

- Casual gamers who enjoy rhythm games (osu!, Beat Saber) and arcade classics
- Players seeking a novel twist on familiar mechanics
- Users with basic web browsers (no downloads or plugins required)

## Core Loop

1. **Listen** — A beat plays (generated via Web Audio API, no external assets needed)
2. **Watch** — Enemies appear and move in sync with the beat
3. **Act** — Player shoots on the beat for bonus points; off-beat shots are less effective
4. **Progress** — Survive waves, earn high scores, unlock new visual themes

## Differentiating Behavior

- **Beat-synced enemy spawning**: Enemies appear exactly on beat intervals
- **On-beat bonus**: Shooting within a ±100ms window of the beat deals 2× damage and awards bonus points
- **Visual pulse**: The screen subtly pulses on each beat, providing visual rhythm cues
- **Dynamic difficulty**: Beat tempo increases gradually across waves

## Acceptance Criteria

| Qualitative Intent | Observable Criterion |
|---|---|
| 新規性 (Novelty) | Enemies spawn on beat; on-beat shooting gives 2× bonus (verifiable in-game) |
| 楽しい (Fun) | Core loop completes a full wave (5+ enemies) without blocking the UI |
| ポップ (Pop) | Visual feedback (screen pulse, score flash) on each beat |
| シンプル (Simple) | Single page app; no configuration required to play |
| Production-ready | Health endpoint `/healthz` returns 200; server starts without errors |

## Non-Goals (Sprint 1)

- Multiplayer or online leaderboards
- External audio files (all audio generated via Web Audio API)
- Mobile touch controls (keyboard only: arrow keys + spacebar)
- Persistent user accounts or save data

## QA Validation

1. Open `http://localhost:8080/` — game UI renders with title "Rhythm Invaders"
2. Press spacebar — enemies spawn on beat (audible click + visual pulse)
3. Shoot on beat — score increases by bonus amount; off-beat shots show reduced score
4. `/healthz` endpoint returns HTTP 200 with `{"status":"ok"}`
5. Server starts with `go run ./cmd/server` and logs startup message
