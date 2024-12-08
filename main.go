package main

import (
	"log"
  "fmt"
  "math/rand"
	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, msg string) {
  for index, char := range msg {
    screen.SetContent(x + index, y, char, nil, tcell.StyleDefault)
  }
}

func setupCoins(level int) []*Sprite {
  coins := make([]*Sprite, level+2)
  for index := range level+2 {
    coins[index] = NewSprite('O', rand.Intn(20), rand.Intn(20))
  }
  return coins
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

	player := NewSprite('@', 10, 10)

  coins := setupCoins(1)

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
      for index, coin := range coins {
        if coin.X == player.X && coin.Y == player.Y {
          // collect coin
          coinCollectedIndex = index
          score++
        }
      }
      // handle coin collisions
      if coinCollectedIndex > -1 {
        // swap target with last
        coins[coinCollectedIndex] = coins[len(coins)-1]
        // trim off last item
        coins = coins[0:len(coins)-1]

        if len(coins) == 0{
          level++
          coins = setupCoins(level)
        }
      }
    } 

	}
}
