# Testing

## Automated validation

Run the configured package scripts when a JavaScript runtime is available:

```sh
npm --prefix client test
npm --prefix client run build
```

The generated scaffold keeps both commands dependency-free by using syntax checks from `client/package.json`.

## Manual smoke check

- Confirm keyboard controls move the defender with ArrowLeft and ArrowRight.
- Confirm Space flips the gravity lane between Floor and Ceiling.
- Confirm Space can score only when the defender is aligned with the invader and on the same lane.
- Confirm the score display updates after a hit.
- Confirm lives decrement when an invader reaches the bottom of the arena.
- Confirm Restart restores score to 0 and lives to 3.
- Confirm the page remains usable on narrow and wide viewports.

## Product Coverage

- Covers the generated gravity-lane invader game and its primary review path.
- Does not claim Docker, Helm, Kubernetes, or CI execution unless those commands were run separately.
