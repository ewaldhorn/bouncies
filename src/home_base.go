package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type HomeBase struct {
	side                                   int
	health, maxHealth                      int
	ticksTillHealthRegeneration            int
	ticksTillCanMaybeFire                  int
	bouncersAvailable, ticksTillNewBouncer int
	xPos, yPos, radius                     float32
	aimPoint                               Vector2D
	baseColour                             color.RGBA
	antialias                              bool
	attackAngle                            float32
}

// ----------------------------------------------------------------------------
type Vector2D struct {
	x, y float32
}

// pre-calced bouncer offsets
var BouncerOffsets = []Vector2D{{x: -15, y: -5}, {x: -5, y: -5}, {x: 5, y: -5}, {x: 15, y: -5}, {x: -15, y: 5}, {x: -5, y: 5}, {x: 5, y: 5}, {x: 15, y: 5}}

// ----------------------------------------------------------------------------
// Sets up a HomeBase with default values
func (h *HomeBase) init(x, y float32) {
	h.maxHealth = DEFAULT_HOMEBASE_HEALTH
	h.health = h.maxHealth
	h.xPos = x
	h.yPos = y
	h.ticksTillNewBouncer = 1
	h.ticksTillHealthRegeneration = DEFAULT_TICKS_PER_SHIELD_REGEN
	h.attackAngle = -36.0
	h.ticksTillCanMaybeFire = DEFAULT_FIRE_DELAY
}

// ----------------------------------------------------------------------------
// Handles respawning of bouncers and shield regeneration
func (h *HomeBase) Update() {
	h.ticksTillNewBouncer -= 1
	if h.ticksTillNewBouncer <= 0 {
		if h.bouncersAvailable < DEFAULT_MAX_BOUNCERS {
			// spawn a new bouncer, ready to be deployed by the player
			h.bouncersAvailable += 1
			h.ticksTillNewBouncer = DEFAULT_TICKS_PER_BOUNCER_RESPAWN
		} else {
			// can't spawn yet, wait
			h.ticksTillNewBouncer = 1
		}
	}

	h.ticksTillHealthRegeneration -= 1
	if h.ticksTillHealthRegeneration <= 0 {
		h.ticksTillHealthRegeneration = DEFAULT_TICKS_PER_SHIELD_REGEN
		if h.health < (h.maxHealth - 5) {
			h.health += 5
		} else {
			h.health = h.maxHealth
		}
	}

	h.ticksTillCanMaybeFire -= 1
	if h.ticksTillCanMaybeFire <= 0 {
		h.ticksTillCanMaybeFire = DEFAULT_FIRE_DELAY
	}
}

// ----------------------------------------------------------------------------
// Allows the HomeBase to take damage, health can be a minimum of 0.
func (h *HomeBase) TakeDamage(amount int) {
	if h.health >= amount {
		h.health -= amount
	} else {
		h.health = 0
	}
}

// ----------------------------------------------------------------------------
// Allows the HomeBase to absorb returning bouncers
func (h *HomeBase) AbsorbShield(amount int) {
	h.health += amount

	if h.health > DEFAULT_HOMEBASE_HEALTH {
		h.health = DEFAULT_HOMEBASE_HEALTH
	}

	if h.bouncersAvailable < DEFAULT_MAX_BOUNCERS {
		h.bouncersAvailable += 1
	}
}

// ----------------------------------------------------------------------------
// Renders the HomeBase on to the provided screen
func (h *HomeBase) Draw(screen *ebiten.Image) {
	var aimX, aimY float32

	// draw attack attackAngle
	if h.side == PLAYER_SIDE {
		aimX = h.xPos + (h.radius+25)*float32(math.Cos(float64(h.attackAngle*math.Pi/180)))
		aimY = h.yPos + (h.radius+25)*float32(math.Sin(float64(h.attackAngle*math.Pi/180)))
	} else {
		aimX = h.xPos - (h.radius+25)*float32(math.Cos(float64(h.attackAngle*math.Pi/180)))
		aimY = h.yPos - (h.radius+25)*float32(math.Sin(float64(h.attackAngle*math.Pi/180)))
	}

	h.aimPoint = Vector2D{x: aimX, y: aimY}

	vector.StrokeLine(screen, h.xPos, h.yPos, aimX, aimY, 3.0, COLOUR_RED, true)

	healthInPercentage := 360 * (float32(h.health*100/DEFAULT_HOMEBASE_HEALTH) / 100)
	radians := healthInPercentage * (math.Pi / 180)

	//fmt.Println("For health at ", h.health, "of", h.maxHealth, "we get", healthInPercentage, "radians", radians)
	// draw shield
	drawArc(screen, h.xPos, h.yPos, h.radius, 5.0, 0.0, radians)

	// now draw base
	vector.DrawFilledCircle(screen, h.xPos, h.yPos, h.radius-1, h.baseColour, h.antialias)

	// finally, draw available bouncers
	for pos := range h.bouncersAvailable {
		vector.DrawFilledCircle(screen, h.xPos+BouncerOffsets[pos].x, h.yPos+BouncerOffsets[pos].y, 4, color.White, h.antialias)
	}

	// debug section
	if IS_DEBUGGING {
		// mid point
		vector.DrawFilledCircle(screen, h.xPos, h.yPos, 5, COLOUR_BLUE, true)
		// aim point
		vector.DrawFilledCircle(screen, aimX, aimY, 5, COLOUR_DARK_RED, true)
		// bounding box
		vector.StrokeRect(screen, h.xPos-h.radius, h.yPos-h.radius, h.radius*2, h.radius*2, 2, COLOUR_BLUE, true)
	}
}

// ----------------------------------------------------------------------------
func (h *HomeBase) AdjustAttackAngle(num float32) {
	if h.side == PLAYER_SIDE {
		h.attackAngle += num
	} else {
		h.attackAngle -= num
	}

	if h.attackAngle < -120.0 {
		h.attackAngle = -120.0
	}

	if h.attackAngle > 28.0 {
		h.attackAngle = 28.0
	}
}

// ----------------------------------------------------------------------------
func (h *HomeBase) AdjustEnemyAttackAngle(direction int) {
	if direction < 10 {
		h.AdjustAttackAngle(-2.0)
	} else if direction > 90 {
		h.AdjustAttackAngle(2.0)
	}
}

// ========================================================== Utility Functions
// Handy for doing odd jobs that are semi-related to the HomeBase struct

// ----------------------------------------------------------------------------
func createPlayerHomeBase() HomeBase {
	playerBase := HomeBase{side: PLAYER_SIDE, radius: 30, baseColour: COLOUR_GREEN, antialias: true}
	playerBase.init(playerBase.radius+DEFAULT_BASE_OFFSET_BUFFER, float32(SCREEN_HEIGHT)-playerBase.radius-DEFAULT_BASE_OFFSET_BUFFER)
	return playerBase
}

// ----------------------------------------------------------------------------
func createEnemyHomeBase() HomeBase {
	enemyBase := HomeBase{side: ENEMY_SIDE, radius: 30, baseColour: COLOUR_RED, antialias: true}
	enemyBase.init(float32(SCREEN_WIDTH)-enemyBase.radius-DEFAULT_BASE_OFFSET_BUFFER, enemyBase.radius+DEFAULT_BASE_OFFSET_BUFFER)
	return enemyBase
}
