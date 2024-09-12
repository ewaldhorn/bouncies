package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type Game struct {
	bases       []HomeBase
	bouncers    []Bouncer
	pressedKeys []ebiten.Key
}

// ----------------------------------------------------------------------------
func (g *Game) initNewGame() {
	g.bases = []HomeBase{createPlayerHomeBase(), createEnemyHomeBase()}
}

// ----------------------------------------------------------------------------
func (g *Game) initBouncers() {
	g.bouncers = []Bouncer{}
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if ebiten.IsFocused() {
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key.String() {
			case "Space":
				b := Bouncer{}
				b.init(PLAYER_SIDE)
				g.bouncers = append(g.bouncers, b)
			}
		}

		for pos, bouncer := range g.bouncers {
			bouncer.update()
			g.bouncers[pos] = bouncer
		}
	}

	for pos, base := range g.bases {
		base.Update()
		g.bases[pos] = base
	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen, 1, 1, float32(SCREEN_WIDTH-1), float32(SCREEN_HEIGHT-1), 0.5, COLOUR_DARK_GRAY, true)
	str := fmt.Sprintf("We are at roughly %.0f FPS, more or less. Focus: %t", ebiten.ActualFPS(), ebiten.IsFocused())
	ebitenutil.DebugPrint(screen, str)

	for i := 0; i < len(g.bouncers); i++ {
		g.bouncers[i].Draw(screen)
	}

	for i := 0; i < len(g.bases); i++ {
		g.bases[i].Draw(screen)
	}
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
