package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var currentObstacleId = 0

// ----------------------------------------------------------------------------
type Obstacle struct {
	id                int
	health, maxHealth int
	xPos, yPos, size  float32
	colour            color.RGBA
}

// ----------------------------------------------------------------------------
// Sets the Obstacle's health to its maximum value.
func (obstacle *Obstacle) initHealth() {
	obstacle.maxHealth = initialHealth
	obstacle.health = obstacle.maxHealth
}

// ----------------------------------------------------------------------------
// Initialises an Obstacle, setting its position, size and health.
func CreateNewObstacle(xPos, yPos, Size float32, colour color.RGBA) *Obstacle {
	obstacle := Obstacle{xPos: xPos, yPos: yPos, size: Size, colour: colour}

	obstacle.id = currentId
	currentId += 1
	if currentId > maxId {
		currentId = 0
	}

	obstacle.initHealth()

	return &obstacle
}

// ----------------------------------------------------------------------------
func (obstacle *Obstacle) TakeHit(num int) {
	obstacle.health -= num
	if obstacle.health < 0 {
		obstacle.health = 0
	} else if obstacle.health > obstacle.maxHealth {
		obstacle.health = obstacle.maxHealth
	}
}

// ----------------------------------------------------------------------------
func (obstacle *Obstacle) Update() {
	if obstacle.health < 30 {
		obstacle.TakeHit(1)
	}
}

// ----------------------------------------------------------------------------
// Renders the Bouncer on to the provided screen
func (obstacle Obstacle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, obstacle.xPos, obstacle.yPos, obstacle.size, obstacle.size, obstacle.colour, true)
}