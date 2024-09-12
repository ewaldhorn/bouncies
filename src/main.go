package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
var FPS int = 1

const (
	GAME_VERSION      = "0.0.1"
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 768
	BOUNCERS      int = 30
)

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Bouncies " + GAME_VERSION)
	ebiten.SetTPS(60)

	game := Game{count: BOUNCERS, lineWidth: 2.0}
	game.initNewGame()
	game.initBouncers()

	err := ebiten.RunGame(&game)

	if err != nil {
		log.Fatal(err)
	}
}
