package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var currentId = 0

// ----------------------------------------------------------------------------
type Bouncer struct {
	side                 int
	id                   int
	health, maxHealth    int
	xPos, yPos, radius   float32
	movementX, movementY float32
	colour               color.RGBA
}

// ----------------------------------------------------------------------------
func (b *Bouncer) init(homeBase HomeBase) {
	b.side = homeBase.side
	b.id = currentId
	currentId += 1
	if currentId > 100000000 {
		currentId = 0
	}

	b.xPos = homeBase.aimPoint.x
	b.yPos = homeBase.aimPoint.y

	b.colour = homeBase.baseColour

	if b.side == PLAYER_SIDE {
		if b.xPos <= homeBase.xPos {
			b.movementX = -1
		} else {
			b.movementX = 1
		}

		if b.yPos <= homeBase.yPos {
			b.movementY = -1
		} else {
			b.movementY = 1
		}
	} else {
		if b.xPos <= homeBase.xPos {
			b.movementX = -1
		} else {
			b.movementX = 1
		}

		if b.yPos <= homeBase.yPos {
			b.movementY = -1
		} else {
			b.movementY = 1
		}
	}

	// TODO: Clean up magic values
	b.radius = 4
	b.maxHealth = 100
	b.health = b.maxHealth
}

// ----------------------------------------------------------------------------
func (b *Bouncer) update() {
	var halfrad = b.radius / 2.0
	b.xPos += b.movementX
	b.yPos += b.movementY

	if b.xPos >= float32(SCREEN_WIDTH-int(halfrad)) || b.xPos <= halfrad {
		b.movementX *= -1
	}
	if b.yPos >= float32(SCREEN_HEIGHT-int(halfrad)) || b.yPos <= halfrad {
		b.movementY *= -1
	}

}

// ----------------------------------------------------------------------------
func (b Bouncer) Draw(screen *ebiten.Image) {
	healthInPercentage := 360 * (float32(b.health*100/b.maxHealth) / 100)
	radians := healthInPercentage * (math.Pi / 180)
	//fmt.Println("For health at ", h.health, "of", h.maxHealth, "we get", healthInPercentage, "radians", radians)

	// first draw shield
	drawArc(screen, b.xPos, b.yPos, b.radius, 3.0, 0.0, radians)

	// now draw bouncer
	vector.DrawFilledCircle(screen, b.xPos, b.yPos, b.radius, b.colour, true)

	if IS_DEBUGGING {
		vector.StrokeRect(screen, b.xPos-b.radius, b.yPos-b.radius, b.radius*2, b.radius*2, 2, COLOUR_BLUE, true)
	}
}

// ----------------------------------------------------------------------------
func (b Bouncer) PrepareVSIS() ([]ebiten.Vertex, []uint16) {
	healthInPercentage := 360 * (float32(b.health*100/b.maxHealth) / 100)
	radians := healthInPercentage * (math.Pi / 180)

	// shield vertices and indices
	vs, is := prepareArcVSIS(b.xPos, b.yPos, b.radius, 0.0, radians)

	// now the actual filled circle too
	ts, ti := prepareCircleVSIS(b.xPos, b.yPos, b.radius, b.colour)
	vs = append(vs, ts...)
	is = append(is, ti...)

	return vs, is
}
