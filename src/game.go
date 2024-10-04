package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
var fontFace = text.NewGoXFace(bitmapfont.Face)

// ----------------------------------------------------------------------------
type Game struct {
	bases       []HomeBase
	bouncers    []Bouncer
	ebitenImage *ebiten.Image
	action      int
	isOver      bool
}

// ----------------------------------------------------------------------------
func (game *Game) initNewGame() {
	game.bases = []HomeBase{createPlayerHomeBase(), createEnemyHomeBase()}
	game.ebitenImage = ebiten.NewImage(SCREEN_WIDTH, SCREEN_HEIGHT)
}

// ----------------------------------------------------------------------------
func (game *Game) initBouncers() {
	game.bouncers = make([]Bouncer, 0, 100)
}

// ----------------------------------------------------------------------------
func (game *Game) updateBouncers() {
	for pos := range game.bouncers {
		game.bouncers[pos].Update()
	}
}

// ----------------------------------------------------------------------------
func (g *Game) updateBases() {
	for pos := range g.bases {
		g.bases[pos].Update()
	}
}

// ----------------------------------------------------------------------------
func (g *Game) checkForGameEnders() {
	if g.bases[PLAYER_SIDE].health <= 5 {
		g.isOver = true
	}

	if g.bases[ENEMY_SIDE].health <= 5 {
		g.isOver = true
	}
}

// ----------------------------------------------------------------------------
func (game *Game) handleMouseInteraction() float64 {
	dx, _ := ebiten.Wheel()

	// check for movement via the wheel
	if dx != 0 {
		angle := math.Copysign(2.0, dx)
		game.bases[PLAYER_SIDE].AdjustAttackAngle(angle)
		game.action = 4
	}

	// check for firing via the mouse
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if didFire, bouncer := game.bases[PLAYER_SIDE].FireBouncer(); didFire {
			game.bouncers = append(game.bouncers, *bouncer)
		}

		game.action = 4
	}

	return dx
}

// ----------------------------------------------------------------------------
func (g *Game) Update() error {
	if ebiten.IsFocused() && !g.isOver {
		dx := g.handleMouseInteraction()
		// handle user interaction

		if g.action <= 0 {

			if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || dx < 0 {
				g.bases[PLAYER_SIDE].AdjustAttackAngle(-2.0)
				g.action = 4
			}

			if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
				g.bases[PLAYER_SIDE].AdjustAttackAngle(-10.0)
				g.action = 4
			}

			if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || dx > 0 {
				g.bases[PLAYER_SIDE].AdjustAttackAngle(2.0)
				g.action = 4
			}

			if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
				g.bases[PLAYER_SIDE].AdjustAttackAngle(10.0)
				g.action = 4
			}

			if ebiten.IsKeyPressed(ebiten.KeySpace) {
				if didFire, bouncer := g.bases[PLAYER_SIDE].FireBouncer(); didFire {
					g.bouncers = append(g.bouncers, *bouncer)
				}
				g.action = 4
			}
		}
		g.action -= 1

		// maybe the enemy feels like firing a shot or six
		if rand.Int()%2 == 0 && (g.bases[ENEMY_SIDE].ticksTillCanMaybeFire <= 1 || g.bases[ENEMY_SIDE].bouncersAvailable >= (DEFAULT_MAX_BOUNCERS-1)) {
			g.bases[ENEMY_SIDE].AdjustEnemyAttackAngle(rand.IntN(100))
			if didFire, bouncer := g.bases[ENEMY_SIDE].FireBouncer(); didFire {
				g.bouncers = append(g.bouncers, *bouncer)
				g.bases[ENEMY_SIDE].ticksTillCanMaybeFire = DEFAULT_FIRE_DELAY
			}

			if g.bases[ENEMY_SIDE].bouncersAvailable >= 2 {
				// maybe try to shoot again soon, because there's some ammo left
				g.bases[ENEMY_SIDE].ticksTillCanMaybeFire = 15
			}
		}

		// now for game object updates
		g.updateBouncers()
		g.updateBases()

		// check if bouncers have hit any bases
		for hpos, base := range g.bases {
			for pos, bouncer := range g.bouncers {
				if bouncer.xPos >= base.centerX-base.radius && bouncer.xPos <= base.centerX+base.radius &&
					bouncer.yPos >= base.centerY-base.radius && bouncer.yPos <= base.centerY+base.radius {
					if base.side == bouncer.side {
						g.bases[hpos].AbsorbShield(bouncer.health)
					} else {
						g.bases[hpos].TakeDamage(bouncer.health)
					}
					g.bouncers[pos].health = 0
				}
			}
		}

		// now check if any bouncers hit any other bouncers
		for outer := 0; outer < len(g.bouncers); outer++ {
			var ob = g.bouncers[outer]
			if ob.health > 0 {
				// only bother if the bouncer has health
				for inner := 0; inner < len(g.bouncers); inner++ {
					if !g.bouncers[outer].hasBounced && g.bouncers[outer].id != g.bouncers[inner].id && g.bouncers[inner].health > 0 {
						var ib = g.bouncers[inner]
						var diff = g.bouncers[inner].radius * 2

						if ob.xPos >= ib.xPos-diff && ob.xPos <= ib.xPos+diff &&
							ob.yPos >= ib.yPos-diff && ob.yPos <= ib.yPos+diff {
							// collided
							g.bouncers[outer].hasBounced = true

							if ob.side != ib.side {
								g.bouncers[outer].TakeHit(15)
								g.bouncers[inner].TakeHit(15)
							} else {
								g.bouncers[outer].movementX *= 1.1
								g.bouncers[outer].movementY *= 1.1
								g.bouncers[inner].movementX *= 1.1
								g.bouncers[inner].movementY *= 1.1

								g.bouncers[outer].TakeHit(-2)
								g.bouncers[inner].TakeHit(-2)
							}

							if rand.Int()%2 == 0 {
								if g.bouncers[outer].movementX > 0 {
									g.bouncers[outer].movementX -= 0.1
								} else {
									g.bouncers[outer].movementX += 0.1
								}
								g.bouncers[outer].movementX *= -1
							}

							if rand.Int()%2 == 0 {
								if g.bouncers[outer].movementY > 0 {
									g.bouncers[outer].movementY -= 0.1
								} else {
									g.bouncers[outer].movementY += 0.1
								}
								g.bouncers[outer].movementY *= -1
							}

							if rand.Int()%2 == 0 {
								if g.bouncers[inner].movementX > 0 {
									g.bouncers[inner].movementX -= 0.1
								} else {
									g.bouncers[inner].movementX += 0.1
								}
								g.bouncers[inner].movementX *= -1
							}

							if rand.Int()%2 == 0 {
								if g.bouncers[inner].movementY > 0 {
									g.bouncers[inner].movementY -= 0.1
								} else {
									g.bouncers[inner].movementY += 0.1
								}
								g.bouncers[inner].movementY *= -1
							}
						}
					}
				}
			}
		}

		// remove dead bouncers
		// TODO optimise, append is horribly slow
		tmpBouncers := make([]Bouncer, 0, 100)
		for _, bouncer := range g.bouncers {
			if bouncer.health >= 12 {
				bouncer.hasBounced = false
				tmpBouncers = append(tmpBouncers, bouncer)
			}
		}
		g.bouncers = tmpBouncers

		g.checkForGameEnders()

	}

	return nil
}

// ----------------------------------------------------------------------------
func (g *Game) Draw(screen *ebiten.Image) {

	g.ebitenImage.Fill(COLOUR_DARK_BLUE)
	vector.StrokeRect(g.ebitenImage, 1, 1, float32(SCREEN_WIDTH-1), float32(SCREEN_HEIGHT-1), 0.5, COLOUR_DARK_GRAY, true)
	str := fmt.Sprintf("(v%s) %.0f FPS vs %.0f TPS. Focus: %t, Angle: %.0f X:%0.f Y:%0.f (%d count)", GAME_VERSION, ebiten.ActualFPS(), ebiten.ActualTPS(), ebiten.IsFocused(), g.bases[0].attackAngle, g.bases[0].aimPoint.x, g.bases[0].aimPoint.y, len(g.bouncers))
	ebitenutil.DebugPrint(g.ebitenImage, str)

	if !g.isOver {
		for i := 0; i < len(g.bouncers); i++ {
			g.bouncers[i].Draw(g.ebitenImage)
		}

		for i := 0; i < len(g.bases); i++ {
			g.bases[i].Draw(g.ebitenImage)
		}

		var ops = &ebiten.DrawImageOptions{}
		screen.DrawImage(g.ebitenImage, ops)
	} else {
		won := g.bases[PLAYER_SIDE].health > 0
		renderGameOverText(screen, won)
	}
}

// ----------------------------------------------------------------------------
func renderGameOverText(screen *ebiten.Image, won bool) {
	textOp := &text.DrawOptions{}

	var str string
	if won {
		str = "Game Over - You Won!"
	} else {
		str = "Gave Over - You Lost!"
	}

	tw, th := text.Measure(str, fontFace, textOp.LineSpacing)

	textOp.GeoM.Translate(float64(SCREEN_WIDTH)/2-(tw/2), float64(SCREEN_HEIGHT)/2-(th/2))
	text.Draw(screen, str, fontFace, textOp)
}

// ----------------------------------------------------------------------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
