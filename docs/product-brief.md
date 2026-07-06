# Product Brief: Gravity Flip Invader

## 1. Concept
**Gravity Flip Invader** is a vertical scrolling shooter (shmup) that subverts the classic "Invader" genre by introducing a gravity inversion mechanic. Instead of being confined to the bottom of the screen, the player can flip gravity to walk on the ceiling, creating a 360-degree tactical playfield within a 2D vertical space.

## 2. Target User
- **Retro Game Enthusiasts**: Players who enjoy classic arcade mechanics but seek a fresh twist.
- **Developers**: Reviewers looking for a clean, idiomatic Go backend with a lightweight frontend, demonstrating clean architecture principles.

## 3. Core Loop
1. **Move**: Player moves left/right and up/down (relative to current gravity).
2. **Flip**: Player presses a key to invert gravity, switching between floor and ceiling.
3. **Shoot**: Player fires bullets towards the "floor" (direction of gravity).
4. **Survive**: Dodge enemy fire and projectiles while eliminating invaders.

## 4. Differentiating Mechanic: Gravity Flip
The core novelty is the **Gravity Flip**. This is not just a visual change; it fundamentally alters gameplay physics and controls.

### Observable Behavior (Review Criteria)
- **Control Inversion**: When gravity is inverted, the "Up" key moves the player towards the ceiling, and "Down" moves them towards the floor.
- **Visual Feedback**: The player sprite rotates 180 degrees upon flipping.
- **Projectile Direction**: Bullets always fire in the direction of the current gravity (downwards relative to the player's current orientation).
- **Collision**: Collision detection updates dynamically based on the player's current position relative to the screen boundaries (top or bottom).

## 5. Non-Goals (Sprint 1)
- **Multiplayer**: Single-player only.
- **Complex Levels**: Procedural or fixed level design is out of scope; focus on a single continuous wave.
- **Sound/Music**: Placeholder or no sound for Sprint 1.
- **High-Res Assets**: Use simple geometric shapes or pixel art placeholders.
- **External Services**: No database, no external APIs. Everything runs locally or in a container.

## 6. Sprint 1 Acceptance Criteria
- [ ] **Gravity Flip**: Player can invert gravity, and controls update accordingly.
- [ ] **Game Loop**: The game runs in a loop with enemies spawning and moving.
- [ ] **Shooting**: Player can fire bullets, and they travel in the direction of gravity.
- [ ] **UI**: A simple HTML/JS frontend renders the game canvas and score.
- [ ] **Backend**: Go server serves the frontend assets and handles game state logic (or serves as a WebSocket/HTTP endpoint for game state if applicable, but for Sprint 1, a static server with client-side logic is acceptable if it demonstrates the architecture).
- [ ] **Docker**: The application can be built and run via Docker.
- [ ] **Helm**: A Helm chart is provided for Kubernetes deployment.
- [ ] **CI**: GitHub Actions runs tests and linting.

## 7. Technical Stack
- **Backend**: Go (net/http, standard library).
- **Frontend**: Vanilla JavaScript, HTML5 Canvas, CSS.
- **Deployment**: Docker, Helm.
- **CI**: GitHub Actions.
