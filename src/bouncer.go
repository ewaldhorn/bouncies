package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/rand"
)

// ----------------------------------------------------------------------------
type Bouncer struct {
	side                 int
	health, maxHealth    int
	xPos, yPos, radius   float32
	movementX, movementY float32
	colour               color.RGBA
}

// ----------------------------------------------------------------------------
func (b *Bouncer) init(side int) {
	b.side = side
	b.xPos = float32(rand.Intn(SCREEN_WIDTH))
	b.yPos = float32(rand.Intn(SCREEN_HEIGHT))

	if rand.Int()%2 == 0 {
		b.movementX = 1
		b.movementY = -1
	} else {
		b.movementX = -1
		b.movementY = 1
	}

	if b.side == PLAYER_SIDE {
		b.colour = COLOUR_GREEN
	} else {
		b.colour = COLOUR_RED
	}

	// TODO: Clean up magic values
	b.radius = 8
	b.maxHealth = 100
	b.health = b.maxHealth
}

// ----------------------------------------------------------------------------
func (b *Bouncer) update() {
	b.xPos += b.movementX
	b.yPos += b.movementY

	if b.xPos >= float32(SCREEN_WIDTH-int(b.radius)) || b.xPos <= b.radius {
		b.movementX *= -1
	}
	if b.yPos >= float32(SCREEN_HEIGHT-int(b.radius)) || b.yPos <= b.radius {
		b.movementY *= -1
	}
}

// ----------------------------------------------------------------------------
func (b Bouncer) Draw(screen *ebiten.Image) {
	healthInPercentage := 360 * (float32(b.health*100/b.maxHealth) / 100)
	radians := healthInPercentage * (math.Pi / 180)
	//fmt.Println("For health at ", h.health, "of", h.maxHealth, "we get", healthInPercentage, "radians", radians)

	// first draw shield
	drawArc(screen, float32(b.xPos), b.yPos, b.radius, 0.0, radians)

	// now draw bouncer
	vector.DrawFilledCircle(screen, b.xPos, b.yPos, b.radius-1, b.colour, true)

}
