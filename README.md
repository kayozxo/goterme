# Coin Collector Game

Coin Collector is a terminal-based game where the player collects coins to progress through levels while avoiding the hidden bomb coin. If the player collects the bomb coin and it’s the last coin, the level progresses; otherwise, the game ends.

---

## Features

- **Dynamic Levels**: Levels increase in difficulty as more coins are added.
- **Bomb Coin Challenge**: One coin is secretly a bomb. Collecting it prematurely ends the game!
- **Intuitive Controls**: Move the player character (`@`) using `W`, `A`, `S`, `D` keys.
- **Game Over Handling**: Displays a final score and exit prompt upon losing.
- **Level Progression**: Successfully collecting all coins, including the bomb as the last coin, advances to the next level.

---

## Demo

### Start Screen

```
Welcome to Coin Collector!
Press 'S' to start the game
Press 'Q' to quit
```

### In-Game

- The player character: `@`
- Coins: `O`
- Bomb Coin: Looks identical to regular coins.

### Game Over

```
Game Over! You hit a bomb!
Final Score: <score>
Press any key to exit
```

---

## Controls

| Key | Action        |
| --- | ------------- |
| `W` | Move Up       |
| `A` | Move Left     |
| `S` | Move Down     |
| `D` | Move Right    |
| `Q` | Quit the Game |

---

## Installation

### Prerequisites

- Go programming language installed
- Terminal that supports ANSI escape codes
- `tcell` library

### Steps

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd coin-collector
   ```

2. Install dependencies:

   ```bash
   go get github.com/gdamore/tcell/v2
   ```

3. Run the game:
   ```bash
   go run main.go
   ```

---

## Code Structure

- **`main.go`**: Contains the game logic, rendering, and input handling.
- **Key Functions**:
  - `setupCoins(level int)`: Sets up coins for the current level and designates a bomb coin.
  - `drawString(screen, x, y, msg)`: Draws strings to the terminal.
  - `startScreen(screen)`: Displays the welcome screen.
  - Game loop handles player movements, coin collection, and level progression.

---

## Gameplay Rules

1. **Objective**: Collect all coins to progress to the next level.
2. **Bomb Coin**:
   - Appears identical to regular coins.
   - If collected prematurely (while other coins remain), ends the game.
   - If collected last, progresses to the next level.
3. **Scoring**: Gain one point for every coin collected.

---

## Future Enhancements

- [] Add a timer for each level.
- [] Introduce obstacles or moving enemies.
- [] Display a leaderboard for high scores.
- [] Add sound effects or animations.

---

## Contribution

Feel free to fork this repository, submit issues, or create pull requests for improvements and bug fixes.

---

## License

This project is licensed under the MIT License. See `LICENSE` for details.
