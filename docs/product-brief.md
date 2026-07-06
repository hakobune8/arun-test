# Product Brief: Product Brief Md

## Concept

Product Brief Md is a compact browser invader game built around a gravity-lane flip mechanic. The player moves left and right, then uses Space to shift between the floor and ceiling lanes. Scoring requires both horizontal alignment and matching the invader lane, so the differentiating mechanic is present in the implemented UI and source code rather than only in documentation.

## Target User

- Players who want a short arcade loop with one clear twist.
- Reviewers who need a fresh-checkout slice that runs without external services.

## Acceptance Criteria

- The visible title, README H1, and this product brief use the same product name.
- The primary route `/` serves the browser game from `client/` when run through the Go server in `server/`.
- Space changes the gravity lane between Floor and Ceiling.
- A score is awarded only when the defender is aligned with the invader and on the same lane.
- The Docker runtime image includes the client assets required for `/` to serve the same UI as local `cd server && go run .`.

## Non-Goals

- Multiplayer.
- External score services.
- Complex level progression.
