package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
var FPS int = 1

const (
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 768
	BOUNCERS      int = 30
)

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Mapper Experiment")
	ebiten.SetTPS(60)

	game := Game{count: BOUNCERS, lineWidth: 2.0}
	game.initNewGame()
	game.initBouncers()

	err := ebiten.RunGame(&game)

	if err != nil {
		log.Fatal(err)
	}
}
