# Goterme

## Overview

This is a simple terminal-based coin collection game developed in Go, utilizing the `tcell` package for creating an interactive terminal user interface. Players navigate a character through a grid, collecting coins and progressing through increasingly challenging levels.

## Features

- Terminal-based gameplay using `tcell` for screen rendering
- Player movement with WASD keys
- Dynamic coin generation
- Scoring system
- Increasing difficulty with each level
- Simple, minimalist design

## Prerequisites

- Go (version 1.16 or newer)
- `tcell` package (`github.com/gdamore/tcell/v2`)

## Installation

1. Ensure Go is installed on your system
2. Clone the repository
3. Install dependencies:

```bash
go mod tidy
go get github.com/gdamore/tcell/v2
```

## How to Play

### Controls

- `W`: Move Up
- `A`: Move Left
- `S`: Move Down
- `D`: Move Right
- `Q`: Quit the game

### Objective

- Collect all coins on the screen to advance to the next level
- Each collected coin increases your score
- The number of coins increases with each level

## Game Mechanics

### Sprite System

The game uses a custom `Sprite` system to represent game entities:

- Player is represented by `@`
- Coins are represented by `O`
- Sprites have X and Y coordinates for positioning

### Level Progression

- Initial game starts at Level 1
- Collecting all coins advances to the next level
- Each new level generates more coins at random positions

### Scoring

- Score increases by 1 for each coin collected
- Current score and level are displayed on the screen

## Code Structure

### Key Components

- `main.go`: Primary game logic
- `setupCoins()`: Generates coins for each level
- `drawString()`: Renders text on the screen
- Game loop handles:
  1. Screen rendering
  2. Player movement
  3. Coin collection
  4. Level progression

## Potential Improvements

- Add obstacles
- Implement player lives
- Create more complex level generation
- Add color and visual enhancements
- Implement a high score system

## Dependencies

- [tcell](https://github.com/gdamore/tcell): Terminal screen handling library

## Running the Game

```bash
go run main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (git checkout -b feature/AmazingFeature)
3. Commit your changes (git commit -m 'Add some AmazingFeature')
4. Push to the branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
