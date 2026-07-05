# Nova Invaders

**Nova Invaders** is a modern, browser-based Invader game built with Go and HTML5 Canvas. It reimagines the classic arcade experience with a core **Phase Shift** mechanic, allowing players to toggle between dimensions to dodge enemy fire and exploit environmental hazards.

## Product Concept & Novelty

*   **Target User:** Retro gaming enthusiasts and players looking for a quick, skill-based browser game.
*   **Core Loop:** Pilot a ship, dodge incoming fire, and destroy waves of enemies.
*   **Differentiating Mechanic:** **Phase Shift**. Players can toggle a "phase" state. In Phase A, enemies are solid but harmless; in Phase B, enemies are intangible but deal damage on contact. This requires strategic timing rather than just reflexes.
*   **Non-Goals:** Multiplayer support, persistent leaderboards, complex narrative.

## Acceptance Criteria

1.  **Playability:** The game is fully playable in modern browsers (Chrome, Firefox, Safari) via the root path `/`.
2.  **Performance:** The Go server handles static assets efficiently; the game loop runs at a stable 60 FPS.
3.  **Health:** The `/health` endpoint returns `200 OK` to verify service availability.
4.  **Deployment:** The application can be containerized with Docker and deployed via Kubernetes (Helm).

## Local Development

### Prerequisites
*   Go 1.21+
*   Docker (optional, for containerized runs)

### Run Locally
```bash
# Start the server
go run ./cmd/server

# Open browser
open http://localhost:8080
```

### Run with Docker
```bash
docker build -t nova-invaders .
docker run -p 8080:8080 nova-invaders
```

## Repository Structure

*   `cmd/server/`: Main application entry point.
*   `internal/game/`: Game logic and state management.
*   `web/`: Frontend assets (HTML, CSS, JS).
*   `charts/`: Helm chart for Kubernetes deployment.
*   `docs/`: Focused documentation for deployment and operations.

## Known Limitations

*   Sound effects are not yet implemented.
*   Mobile touch controls are in early prototype stage.

## Next Steps

*   Implement sound effects and music.
*   Add particle effects for explosions.
*   Optimize asset loading for mobile devices.
