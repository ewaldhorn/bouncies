package main

import (
	"image/color"

	"golang.org/x/exp/rand"
)

// ----------------------------------------------------------------------------
type Bouncer struct {
	positionX, positionY float64
	movementX, movementY float64
	colour               color.RGBA
}

// ----------------------------------------------------------------------------
func (b *Bouncer) init() {
	b.positionX = float64(rand.Intn(SCREEN_WIDTH))
	b.positionY = float64(rand.Intn(SCREEN_HEIGHT))

	if rand.Int()%2 == 0 {
		b.movementX = 1
		b.movementY = -1
	} else {
		b.movementX = -1
		b.movementY = 1
	}

	b.colour = color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
}

// ----------------------------------------------------------------------------
func (b *Bouncer) update() {
	b.positionX += b.movementX
	b.positionY += b.movementY

	if b.positionX >= float64(SCREEN_WIDTH-2) || b.positionX <= 2 {
		b.movementX *= -1
	}
	if b.positionY >= float64(SCREEN_HEIGHT-2) || b.positionY <= 2 {
		b.movementY *= -1
	}
}
