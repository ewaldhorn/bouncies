package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ----------------------------------------------------------------------------
var TPS int = 30

const (
	GAME_VERSION      = "0.0.4b"
	IS_DEBUGGING      = false
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 768
)

// ----------------------------------------------------------------------------
func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Bouncies " + GAME_VERSION)
	ebiten.SetTPS(TPS)

	game := Game{}
	game.initNewGame()
	game.initBouncers()

	err := ebiten.RunGame(&game)

	if err != nil {
		log.Fatal(err)
	}
}
