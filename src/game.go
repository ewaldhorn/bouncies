package main

import (
	"fmt"
	"math/rand/v2"

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
		// handle user interaction
		g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])
		for _, key := range g.pressedKeys {
			switch key.String() {
			case "ArrowUp":
				g.bases[0].AdjustAttackAngle(-4.0)
				g.bases[1].AdjustAttackAngle(-4.0)
			case "ArrowDown":
				g.bases[0].AdjustAttackAngle(4.0)
				g.bases[1].AdjustAttackAngle(4.0)
			case "Space":
				if g.bases[PLAYER_SIDE].bouncersAvailable > 0 {
					g.bases[PLAYER_SIDE].bouncersAvailable -= 1
					b := Bouncer{}
					b.init(g.bases[PLAYER_SIDE])
					g.bouncers = append(g.bouncers, b)
				}
			}
		}

		// maybe the enemy feels like firing a shot or six
		if g.bases[ENEMY_SIDE].ticksTillCanMaybeFire <= 1 {
			if rand.Int()%2 == 0 {
				if g.bases[ENEMY_SIDE].bouncersAvailable > 0 {
					for count := 0; count < rand.IntN(g.bases[ENEMY_SIDE].bouncersAvailable); count++ {
						g.bases[ENEMY_SIDE].bouncersAvailable -= 1
						b := Bouncer{}
						b.init(g.bases[ENEMY_SIDE])
						g.bouncers = append(g.bouncers, b)
					}
				}
			}
		}

		// now for game object updates
		for pos, bouncer := range g.bouncers {
			bouncer.update()
			g.bouncers[pos] = bouncer
		}

		for pos, base := range g.bases {
			base.Update()
			g.bases[pos] = base
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen, 1, 1, float32(SCREEN_WIDTH-1), float32(SCREEN_HEIGHT-1), 0.5, COLOUR_DARK_GRAY, true)
	str := fmt.Sprintf("We are at roughly %.0f FPS, more or less. Focus: %t, Angle: %.0f X:%0.f Y:%0.f (%d count)", ebiten.ActualFPS(), ebiten.IsFocused(), g.bases[0].attackAngle, g.bases[0].aimPoint.x, g.bases[0].aimPoint.y, len(g.bouncers))
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
