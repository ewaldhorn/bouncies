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
	count       int
	bases       []HomeBase
	bouncers    []Bouncer
	pressedKeys []ebiten.Key
	lineWidth   float32
}

// ----------------------------------------------------------------------------
func (g *Game) initNewGame() {
	g.bases = make([]HomeBase, 0)

	playerBase := HomeBase{radius: 30, baseColour: COLOUR_GREEN, antialias: true}
	playerBase.init(playerBase.radius+DEFAULT_BASE_OFFSET_BUFFER, float32(SCREEN_HEIGHT)-playerBase.radius-DEFAULT_BASE_OFFSET_BUFFER)

	enemyBase := HomeBase{radius: 30, baseColour: COLOUR_RED, antialias: true}
	enemyBase.init(float32(SCREEN_WIDTH)-enemyBase.radius-DEFAULT_BASE_OFFSET_BUFFER, enemyBase.radius+DEFAULT_BASE_OFFSET_BUFFER)

	g.bases = append(g.bases, playerBase, enemyBase)
}

// ----------------------------------------------------------------------------
func (g *Game) initBouncers() {
	g.bouncers = make([]Bouncer, g.count)

	for bouncerPosition := range g.bouncers {
		tmpBouncer := Bouncer{}
		tmpBouncer.init()
		g.bouncers[bouncerPosition] = tmpBouncer
	}
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if ebiten.IsFocused() {
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

		for _, key := range g.pressedKeys {
			switch key.String() {
			case "ArrowDown":
				if g.lineWidth > 0.20 {
					g.lineWidth -= 0.10
				}
			case "ArrowUp":
				if g.lineWidth < 50.0 {
					g.lineWidth += 0.10
				}
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
	str := fmt.Sprintf("We are at roughly %.0f FPS, more or less. (Line: %0.2f) Focus: %t", ebiten.ActualFPS(), g.lineWidth, ebiten.IsFocused())
	ebitenutil.DebugPrint(screen, str)

	for i := 1; i < len(g.bouncers); i++ {
		vector.StrokeLine(screen,
			float32(g.bouncers[i-1].positionX),
			float32(g.bouncers[i-1].positionY),
			float32(g.bouncers[i].positionX),
			float32(g.bouncers[i].positionY),
			g.lineWidth,
			g.bouncers[i].colour, true)
	}

	lastBouncer := g.bouncers[len(g.bouncers)-1]
	vector.StrokeLine(screen,
		float32(lastBouncer.positionX), float32(lastBouncer.positionY),
		float32(g.bouncers[0].positionX), float32(g.bouncers[0].positionY),
		g.lineWidth, lastBouncer.colour, true)

	for i := 0; i < len(g.bases); i++ {
		g.bases[i].Draw(screen)
	}
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
