# Product/Design Brief: Neon Orbit Invaders

## 1. Product Concept
**Neon Orbit Invaders** is a modern twist on the classic arcade shooter. Instead of linear side-to-side movement, enemies orbit a central planetary core. The player controls a ship that can rotate 360 degrees to aim and fire. The core novelty lies in **Orbital Timing**: enemies have a "vulnerable window" as they pass a specific point in their orbit. Hitting them during this window deals critical damage, rewarding precision over spam.

## 2. Target User
- Casual arcade gamers who enjoy retro aesthetics but want fresh mechanics.
- Players looking for a quick, satisfying loop with clear visual feedback.

## 3. Core Loop / Workflow
1. **Spawn**: Enemies spawn and begin orbiting the central core at varying speeds and radii.
2. **Aim & Fire**: Player moves the ship and rotates the barrel to aim. Fires projectiles.
3. **Orbital Timing**: Player tracks enemy orbits and fires during the "vulnerable window" for instant destruction.
4. **Feedback**: Visual "crunch" effect and score update on hit. Misses require re-engaging.
5. **Progression**: Wave completion increases enemy count and orbital complexity.

## 4. Differentiating Behavior
- **Orbital Movement**: Enemies follow circular/elliptical paths, not linear grids.
- **360-Degree Aiming**: Player ship rotates to aim, not just move left/right.
- **Vulnerable Window**: Visual indicator (e.g., color shift or glow) shows when an enemy is vulnerable, adding a rhythm/timing layer.

## 5. Non-Goals
- Multiplayer or online leaderboards (Sprint 1).
- Complex power-ups or inventory systems (Sprint 1).
- Persistent user accounts or cloud saves (Sprint 1).
- Mobile touch controls (Sprint 1 focuses on keyboard/mouse).

## 6. Acceptance Criteria
- [ ] Game renders in browser with 60 FPS target.
- [ ] Enemies orbit a central core with visible vulnerable windows.
- [ ] Player can aim 360 degrees and fire projectiles.
- [ ] Hitting an enemy during the vulnerable window destroys it instantly.
- [ ] Score updates in real-time.
- [ ] Health endpoint returns 200 OK.
- [ ] Docker image builds successfully.
- [ ] Helm chart deploys to Kubernetes without errors.

## 7. Qualitative Requirements → Observable Criteria
| Requirement | Observable Criteria |
|-------------|---------------------|
| **Novelty** | Enemies move in orbits, not linear paths. Aiming is 360-degree, not fixed upward. |
| **Fun** | Satisfying visual feedback on hit (particle effect). Increasing difficulty curve. |
| **Pop** | High-contrast neon colors on dark background. Simple geometric shapes. |
| **Simple** | One action to shoot. Clear vulnerable window indicator. |
| **Production-Ready** | Health endpoint, Dockerfile, Helm chart, CI pipeline, structured logs. |
