package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, msg string) {
  for index, char := range msg {
    screen.SetContent(x + index, y, char, nil, tcell.StyleDefault)
  }
}

func setupCoins(level int) ([]*Sprite, int) {
  coins := make([]*Sprite, level+2)
  for i := range coins {
    coins[i] = NewSprite('O', rand.Intn(20), rand.Intn(20)) // All coins look the same
  }
  bombIndex := rand.Intn(len(coins)) // Randomly pick the bomb coin
  return coins, bombIndex
}

func startScreen(screen tcell.Screen) {
  screen.Clear()
  drawString(screen, 10, 10, "Welcome to Coin Collector!")
  drawString(screen, 10, 12, "Press 'S' to start the game")
  drawString(screen, 10, 14, "Press 'Q' to quit")
  screen.Show()

  run := true
  for run {
    ev := screen.PollEvent()
    switch ev := ev.(type) {
      case *tcell.EventKey:
        switch ev.Rune() {
          case 's':
            return
          case 'q':
            screen.Fini()
            run = false
        }
    }
  }
}

func main () {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	//game init section
  startScreen(screen)

	player := NewSprite('@', 10, 10)

  coins, bombIndex := setupCoins(1)

  score := 0
  level := 1

	// game loop
	running := true
	for running {
		//draw logic
		screen.Clear()

		player.Draw(screen)

    for _, coin := range coins {
      coin.Draw(screen)
    }

    // ui
    drawString(screen, 1, 1, fmt.Sprintf("Score: %d", score))
    drawString(screen, 1, 2, fmt.Sprintf("Level: %d", level))

		screen.Show()
		//update logic

    playerMoved := false

		// getting the event
		ev := screen.PollEvent()
		//checking event type
		switch ev := ev.(type) {
			case *tcell.EventKey:
				//checking event key
				switch ev.Rune() {
					case 'q':
						running = false
					case 'w':
						player.Y -= 1
            playerMoved = true
					case 'a':
						player.X -= 1
            playerMoved = true
					case 's':
						player.Y += 1
            playerMoved = true
					case 'd':
						player.X += 1
            playerMoved = true
				}
		}

    // check for coin collisions
    if playerMoved {
      coinCollectedIndex := -1
      isBomb := false

      // Check if the player collects a coin
      for index, coin := range coins {
        if coin.X == player.X && coin.Y == player.Y {
          coinCollectedIndex = index
          isBomb = (index == bombIndex)
          break
        }
      }

      if coinCollectedIndex > -1 {
        if isBomb {
          if len(coins) == 1 {
            // Bomb is the last coin; show message and progress to next level
            screen.Clear()
            drawString(screen, 10, 10, "You found the bomb, but completed the level!")
            drawString(screen, 10, 12, "Press any key to continue to next level")
            screen.Show()
            screen.PollEvent() // Wait for key press

            // Progress to next level
            level++
            coins, bombIndex = setupCoins(level)
            score++
          } else {
            // Bomb collected and other coins remain: Game Over
            screen.Clear()
            drawString(screen, 10, 10, "Game Over! You hit a bomb!")
            drawString(screen, 10, 12, fmt.Sprintf("Final Score: %d", score))
            drawString(screen, 10, 14, "Press any key to exit")
            screen.Show()
            screen.PollEvent() // Wait for user input before exiting
            return
          }
        } else {
          // Regular coin collected
          score++
          // Remove the collected coin
          coins[coinCollectedIndex] = coins[len(coins)-1]
          coins = coins[:len(coins)-1]

          // If the bomb was the last coin in the array, update its index
          if bombIndex == len(coins) {
            bombIndex = coinCollectedIndex
          }

          // Check if all coins are collected
          if len(coins) == 0 {
            level++
            coins, bombIndex = setupCoins(level)
          }
        }
      }
    }
  }
}