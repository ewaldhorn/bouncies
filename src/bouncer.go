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
	speedup              int
	age                  int
	hasBounced           bool
	health, maxHealth    int
	xPos, yPos, radius   float32
	movementX, movementY float32
	shieldRadians        float32
	colour               color.RGBA
}

// ----------------------------------------------------------------------------
func (b *Bouncer) Init(homeBase HomeBase) {
	b.side = homeBase.side
	b.id = currentId
	currentId += 1
	if currentId > 100000000 {
		currentId = 0
	}

	b.hasBounced = false
	b.xPos = homeBase.aimPoint.x
	b.yPos = homeBase.aimPoint.y

	b.colour = homeBase.baseColour

	aimX := homeBase.xPos + (homeBase.radius+25)*float32(math.Cos(float64(homeBase.attackAngle*math.Pi/180)))
	aimY := homeBase.yPos + (homeBase.radius+25)*float32(math.Sin(float64(homeBase.attackAngle*math.Pi/180)))

	deltaX := aimX - homeBase.xPos
	deltaY := aimY - homeBase.yPos

	angle := math.Atan2(float64(deltaY), float64(deltaX))
	b.movementX = float32(math.Cos(angle)) * 1.5
	b.movementY = float32(math.Sin(angle)) * 1.5

	if b.side == ENEMY_SIDE {
		b.movementX *= -1.0
		b.movementY *= -1.0
	}

	// TODO: Clean up magic values
	b.radius = 4
	b.maxHealth = 100
	b.health = b.maxHealth

	// need to set the shield data now
	b.updateShield()
}

// ----------------------------------------------------------------------------
func (b *Bouncer) TakeHit(num int) {
	b.health -= num
	if b.health < 0 {
		b.health = 0
	} else if b.health > b.maxHealth {
		b.health = b.maxHealth
	}

	b.updateShield()
}

// ----------------------------------------------------------------------------
func (b *Bouncer) updateShield() {
	healthInPercentage := 360 * (float32(b.health*100/b.maxHealth) / 100)
	b.shieldRadians = healthInPercentage * RADIAN
}

// ----------------------------------------------------------------------------
// Bouncers gain energy from bouncing of the sides
func (b *Bouncer) Update() {
	b.age += 1

	var halfrad = b.radius / 2.0
	b.xPos += b.movementX
	b.yPos += b.movementY

	if b.xPos >= float32(SCREEN_WIDTH-int(halfrad)) || b.xPos <= halfrad {
		b.movementX *= -1.2
		b.TakeHit(5)
	}

	if b.yPos >= float32(SCREEN_HEIGHT-int(halfrad)) || b.yPos <= halfrad {
		b.movementY *= -1.2
		b.TakeHit(5)
	}

	if math.Abs(float64(b.movementX)) <= 0.1 && math.Abs(float64(b.movementY)) <= 0.1 {
		b.TakeHit(2)
	}

	b.speedup += 1

	if b.speedup > 30 {
		b.speedup = 0
		b.movementX *= 1.1
		b.movementY *= 1.1
	}

	if b.movementX > 5.0 {
		b.movementX = 5.0
	}

	if b.movementX < -5.0 {
		b.movementX = -5.0
	}

	if b.movementY > 5.0 {
		b.movementY = 5.0
	}

	if b.movementY < -5.0 {
		b.movementY = -5.0
	}

	if b.health < 30 || b.age > 1000 {
		b.TakeHit(1)
	}
}

// ----------------------------------------------------------------------------
func (b Bouncer) Draw(screen *ebiten.Image) {
	// first draw shield
	drawFilledArc(screen, b.xPos, b.yPos, b.radius+2, 0.0, b.shieldRadians, color.White)

	// now draw bouncer
	vector.DrawFilledCircle(screen, b.xPos, b.yPos, b.radius, b.colour, true)

	if IS_DEBUGGING {
		vector.StrokeRect(screen, b.xPos-b.radius, b.yPos-b.radius, b.radius*2, b.radius*2, 2, COLOUR_BLUE, true)
	}
}
