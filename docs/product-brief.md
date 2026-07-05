# Product Brief: Gravity Invaders

## 1. Game Concept
**Gravity Invaders** is a minimalist, high-contrast arcade shooter where the player controls a ship that can instantly flip its gravity. Instead of traditional movement, the ship falls up or down, bouncing off the top and bottom of the screen while shooting horizontally at waves of invaders. The core loop emphasizes rhythm, timing, and spatial awareness over reflex-heavy dodging.

## 2. Target User
- Arcade enthusiasts who enjoy fast-paced, minimalist games.
- Players looking for a "pick up and play" experience with high replayability.
- Users who appreciate clean, retro-inspired aesthetics with modern mechanics.

## 3. Core Loop / Workflow
1. **Start**: Player selects a ship and begins the level.
2. **Flip**: Player presses a key to flip gravity (Up/Down).
3. **Shoot**: Player fires projectiles horizontally at invaders.
4. **Survive**: Avoid collisions with enemy projectiles and the invaders themselves.
5. **Progress**: Clear the wave to advance to the next level with increased difficulty.

## 4. Differentiating Mechanic: Gravity Flip
- **Mechanic**: The ship does not move vertically via continuous input. Instead, a single button press flips the ship's gravity direction.
- **Impact**: This creates a unique movement pattern where the ship arcs across the screen, requiring players to anticipate trajectories rather than react to them.
- **Validation**: The mechanic must feel responsive and intuitive. The flip should be instant, with no delay or animation lag.

## 5. Acceptance Criteria
- [ ] **Gravity Flip**: The ship flips gravity direction instantly upon button press.
- [ ] **Shooting**: The ship fires projectiles horizontally in the direction it is facing.
- [ ] **Invaders**: Invaders move in a classic pattern, decreasing in speed as fewer remain.
- [ ] **Collision**: The game detects collisions between projectiles and invaders, and between the ship and invaders/projectiles.
- [ ] **Game Over**: The game ends when the ship is destroyed or invaders reach the bottom.
- [ ] **Score**: A score counter updates when invaders are destroyed.
- [ ] **UI**: The game displays a start screen, game over screen, and current score.

## 6. Non-Goals
- Multiplayer functionality.
 - Complex power-ups or item drops.
- Story mode or narrative elements.
- High-fidelity graphics or sound effects (placeholder assets are acceptable for Sprint 1).
- Mobile responsiveness (desktop browser focus for Sprint 1).

## 7. Qualitative Requirements -> Observable Criteria
- **Novelty**: The gravity flip mechanic is unique compared to traditional invader games. *QA Check*: Verify that the ship moves vertically only via gravity flips, not direct control.
- **Fun**: The game is engaging and encourages replay. *QA Check*: Playtest for at least 5 minutes; note if the player feels compelled to continue.
- **Pop/Simple**: The visual style is clean and easy to understand. *QA Check*: Verify that the UI is minimal, with clear contrast between the ship, invaders, and background.
- **Production-Ready**: The game runs smoothly in a modern browser. *QA Check*: Verify that the game runs at 60 FPS on a standard laptop.
