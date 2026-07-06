# Smoke Test

1. Start the Go server with `cd server && go run .` and open `/`, or open `client/index.html` in a browser.
2. Confirm the arena, score display, lives display, gravity display, target lane display, and restart button render without layout overlap.
3. Press ArrowLeft and ArrowRight and confirm the defender moves horizontally.
4. Press Space and confirm the gravity display flips between Floor and Ceiling.
5. Press Space while aligned on the same lane as the invader and confirm the score increases.
6. Let an invader reach the bottom and confirm lives decrement.
7. Click `Restart` and confirm score returns to 0, lives returns to 3, and gravity returns to Floor.
8. Run `npm test` and `npm run build`.

## Product Coverage

- Validates the generated browser game from the primary served route.
- Focuses on implemented controls, scoring, lives, restart behavior, and layout.
