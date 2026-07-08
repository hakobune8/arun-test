---
title: "Gravity Shift Invaders - Product Brief"
status: "active"
sprint: 1
---

# Gravity Shift Invaders

## 1. Concept & Value Proposition

**Gravity Shift Invaders** is a modern twist on the classic arcade shooter. 
Instead of the traditional left-right movement, the player controls the **gravity** of their ship.

*   **Novelty:** Vertical movement as the primary mechanic. The ship automatically moves forward (or the stage scrolls), and the player taps/clicks to flip gravity, making the ship fall to the bottom or float to the top.
*   **Value:** A fresh, intuitive arcade experience that feels different from standard Invaders while retaining the satisfying "shoot and destroy" loop.
*   **Target User:** Casual gamers, arcade enthusiasts, and users looking for a quick, high-score-chasing experience.

## 2. Core Loop

1.  **Start:** The game begins with the ship at the bottom.
2.  **Input:** Player clicks/taps to flip gravity.
    *   If at bottom, ship floats to top.
    *   If at top, ship falls to bottom.
3.  **Action:** The ship automatically fires bullets forward (or player presses space to fire, depending on implementation simplicity for Sprint 1).
4.  **Feedback:** Enemies are destroyed, score increases, visual effects play.
5.  **Progression:** Enemies spawn in waves. Difficulty increases over time.
6.  **End:** Game over when the player's ship is hit by an enemy or enemy bullet.

## 3. Differentiating Mechanics

*   **Gravity Flip:** The core interaction. No left/right movement. This creates a unique "dodging" dynamic where players must anticipate enemy patterns vertically.
*   **Auto-Fire:** To keep controls simple and focused on the gravity mechanic, the ship auto-fires continuously.

## 4. Acceptance Criteria (Sprint 1)

*   [ ] **AC-1: Gravity Control**
    *   Given the game is running,
    *   When the user clicks/taps,
    *   Then the ship changes vertical direction (falls or floats).
*   [ ] **AC-2: Auto-Fire**
    *   Given the game is running,
    *   Then the ship fires bullets automatically at a set interval.
*   [ ] **AC-3: Enemy Movement**
    *   Given enemies are spawned,
    *   Then they move in a recognizable pattern (e.g., straight down or sine wave).
*   [ ] **AC-4: Collision Detection**
    *   Given a bullet hits an enemy,
    *   Then the enemy is removed and score increases.
    *   Given an enemy hits the player,
    *   Then the game ends (or player loses a life).
*   [ ] **AC-5: UI Feedback**
    *   The current score is displayed.
    *   The game state (Playing, Game Over) is clearly visible.

## 5. Non-Goals (Sprint 1)

*   Multiplayer support.
*   Complex power-ups (beyond basic shooting).
*   Persistent high scores (local storage only if easy, otherwise in-memory).
*   Mobile-specific touch gestures (swipe, etc.) - simple tap/click is sufficient for MVP.

## 6. Technical Constraints

*   **Frontend:** Vanilla HTML5 Canvas + JavaScript (no heavy frameworks for Sprint 1).
*   **Backend:** Go HTTP server serving static assets and handling health checks.
*   **Deployment:** Dockerized, deployable via Helm/K8s.
