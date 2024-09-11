package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
type HomeBase struct {
	health, maxHealth  int
	xPos, yPos, radius float32
	baseColour         color.RGBA
	antialias          bool
}

// ----------------------------------------------------------------------------
// Sets up a HomeBase with sane values
func (h *HomeBase) init(x, y float32) {
	h.maxHealth = HOMEBASE_DEFAULT_HEALTH
	h.health = h.maxHealth
	h.xPos = x
	h.yPos = y
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
func (h HomeBase) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, h.xPos, h.yPos, h.radius, h.baseColour, h.antialias)
}

// ----------------------------------------------------------------------------
