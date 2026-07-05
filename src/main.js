const arena = document.querySelector(".arena");
const player = document.getElementById("player");
const invader = document.getElementById("invader");
const scoreEl = document.getElementById("score");
const livesEl = document.getElementById("lives");
const gravityEl = document.getElementById("gravity");
const targetLaneEl = document.getElementById("target-lane");
const statusEl = document.getElementById("status");
const restartButton = document.getElementById("restart");

const state = {
  score: 0,
  lives: 3,
  playerX: 50,
  invaderX: 50,
  invaderY: 8,
  gravity: "floor",
  invaderLane: "floor"
};

function render() {
  player.style.transform = `translateX(calc(${state.playerX}% - 22px))`;
  invader.style.transform = `translate(calc(${state.invaderX}% - 22px), ${state.invaderY}px)`;
  player.classList.toggle("ceiling", state.gravity === "ceiling");
  invader.classList.toggle("ceiling", state.invaderLane === "ceiling");
  invader.style.top = state.invaderLane === "ceiling" ? "auto" : "24px";
  invader.style.bottom = state.invaderLane === "ceiling" ? "24px" : "auto";
  scoreEl.textContent = String(state.score);
  livesEl.textContent = String(state.lives);
  gravityEl.textContent = state.gravity === "ceiling" ? "Ceiling" : "Floor";
  targetLaneEl.textContent = state.invaderLane === "ceiling" ? "Ceiling" : "Floor";
}

function resetInvader() {
  state.invaderX = 15 + ((state.score * 29) % 70);
  state.invaderY = 8;
  state.invaderLane = state.score % 20 === 0 ? "ceiling" : "floor";
}

function restart() {
  state.score = 0;
  state.lives = 3;
  state.playerX = 50;
  state.gravity = "floor";
  resetInvader();
  statusEl.textContent = "Ready. Flip gravity to match the target lane.";
  render();
}

function flipGravity() {
  if (state.lives <= 0) {
    return;
  }
  state.gravity = state.gravity === "floor" ? "ceiling" : "floor";
  const aligned = Math.abs(state.playerX - state.invaderX) <= 12;
  const laneMatched = state.gravity === state.invaderLane;
  if (aligned && laneMatched) {
    state.score += state.gravity === "ceiling" ? 15 : 10;
    statusEl.textContent = "Gravity match. Score increased.";
    resetInvader();
  } else if (!laneMatched) {
    statusEl.textContent = "Wrong lane. Flip gravity to match the target.";
  } else {
    statusEl.textContent = "Lane matched. Align horizontally before the next flip.";
  }
  render();
}

function tick() {
  if (state.lives > 0) {
    state.invaderY += 4;
    if (state.invaderY > arena.clientHeight - 82) {
      state.lives -= 1;
      statusEl.textContent = state.lives > 0 ? "Invader slipped through." : "Game over. Restart to try again.";
      resetInvader();
    }
    render();
  }
  window.setTimeout(tick, 320);
}

document.addEventListener("keydown", (event) => {
  if (event.key === "ArrowLeft") {
    state.playerX = Math.max(8, state.playerX - 8);
  } else if (event.key === "ArrowRight") {
    state.playerX = Math.min(92, state.playerX + 8);
  } else if (event.code === "Space") {
    event.preventDefault();
    flipGravity();
  } else {
    return;
  }
  render();
});

restartButton.addEventListener("click", restart);

restart();
tick();
