package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var currentId = 0

const (
	maxId         = 100000000
	aimOffset     = 25
	bouncerRadius = 4
	initialHealth = 100
)

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
// Sets the Bouncer's initial position to the HomeBase's aim point
func (b *Bouncer) initPosition(homeBase HomeBase) {
	b.xPos = homeBase.aimPoint.x
	b.yPos = homeBase.aimPoint.y
}

// ----------------------------------------------------------------------------
// Sets the Bouncer's initial movement to head towards the enemy base.
// The exact movement is determined by the HomeBase's attackAngle.
func (b *Bouncer) initVelocity(homeBase HomeBase) {
	aimX := homeBase.centerX + (homeBase.radius+aimOffset)*float32(math.Cos(homeBase.attackAngle*math.Pi/180))
	aimY := homeBase.centerY + (homeBase.radius+aimOffset)*float32(math.Sin(homeBase.attackAngle*math.Pi/180))

	horizontalOffset := aimX - homeBase.centerX
	verticalOffset := aimY - homeBase.centerY

	angle := math.Atan2(float64(verticalOffset), float64(horizontalOffset))
	b.movementX = float32(math.Cos(angle) * 1.5)
	b.movementY = float32(math.Sin(angle) * 1.5)

	if b.side == ENEMY_SIDE {
		b.movementX *= -1.0
		b.movementY *= -1.0
	}
}

// ----------------------------------------------------------------------------
// Sets the Bouncer's health to its maximum value.
func (b *Bouncer) initHealth() {
	b.maxHealth = initialHealth
	b.health = b.maxHealth
}

// ----------------------------------------------------------------------------
// Initialises a Bouncer, setting its position, velocity and health. It also
// sets the Bouncer's shield data.
func (b *Bouncer) Init(homeBase HomeBase) {
	b.side = homeBase.side
	b.id = currentId
	currentId += 1
	if currentId > maxId {
		currentId = 0
	}

	b.hasBounced = false
	b.initPosition(homeBase)
	b.initVelocity(homeBase)
	b.initHealth()

	b.colour = homeBase.baseColour
	b.radius = bouncerRadius

	// need to set the shield data now
	b.updateShield()
}

// ----------------------------------------------------------------------------
// TakeHit reduces the Bouncer's health by the given number. It
// will not go below 0, and will not go above the maximum health.
// When the health changes, the shield is updated.
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
// Calculates the shield angle based on the current health. The shield
// angle is the number of radians that the shield arc should cover.
func (b *Bouncer) updateShield() {
	healthInPercentage := 360 * (float32(b.health*100/b.maxHealth) / 100)
	b.shieldRadians = healthInPercentage * RADIAN
}

// ----------------------------------------------------------------------------
// The Update function is called once per frame and moves the Bouncer one step. It
// also checks for collisions with the screen edges and adjusts the Bouncer's
// movement accordingly. The Bouncer also gains speed over time up to a limit.
// The Bouncer's health is decreased over time, and there is a minimum health
// threshold.
func (b *Bouncer) Update() {
	b.age += 1

	b.xPos += b.movementX
	b.yPos += b.movementY

	if b.xPos >= float32(SCREEN_WIDTH-int(b.radius)) || b.xPos <= b.radius {
		b.movementX *= -1.2
		b.TakeHit(5)
	}

	if b.yPos >= float32(SCREEN_HEIGHT-int(b.radius)) || b.yPos <= b.radius {
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
// Renders the Bouncer on to the provided screen
func (b Bouncer) Draw(screen *ebiten.Image) {
	// first draw shield
	drawFilledArc(screen, b.xPos, b.yPos, b.radius+2, 0.0, b.shieldRadians, color.White)

	// now draw bouncer
	vector.DrawFilledCircle(screen, b.xPos, b.yPos, b.radius, b.colour, true)

	if IS_DEBUGGING {
		vector.StrokeRect(screen, b.xPos-b.radius, b.yPos-b.radius, b.radius*2, b.radius*2, 2, COLOUR_BLUE, true)
	}
}
