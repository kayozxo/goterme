package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

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

	// game loop
	running := true
	for running {
		//draw logic

		screen.Clear()

		player.Draw(screen)

		screen.Show()
		//update logic

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
					case 'a':
						player.X -= 1
					case 's':
						player.Y += 1
					case 'd':
						player.X += 1
				}
		}
	}
}
