# One-Button Invaders

This repository started empty. ARUN generated a minimal static browser game with a gravity-lane mechanic so an implementation-heavy scrum workflow can produce reviewable code, documentation, and validation artifacts without GitHub API calls.

## Features

- Keyboard controls with ArrowLeft, ArrowRight, and Space.
- Space flips the defender between floor and ceiling gravity lanes.
- Score display that increments only when the defender is horizontally aligned and on the same gravity lane as the invader.
- Lives tracking that decrements when an invader reaches the bottom of the arena.
- Restart behavior that resets score, lives, player position, and invader position.

## Run

Open `index.html` in a browser, or serve the directory with any static file server.

## Validate

```sh
npm test
npm run build
```

Both scripts use `node --check` and do not require package installation.
