package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var currentObstacleId = 0

const obstacleInitialHealth = 400
const obstacleSize = 32

// ----------------------------------------------------------------------------
type Obstacle struct {
	id                int
	nextMove          int
	health, maxHealth int
	healthTimer       int
	xPos, yPos, size  float32
	colour            color.RGBA
}

// ----------------------------------------------------------------------------
// Sets the Obstacle's health to its maximum value.
func (obstacle *Obstacle) initHealth() {
	obstacle.maxHealth = obstacleInitialHealth
	obstacle.health = obstacle.maxHealth
}

// ----------------------------------------------------------------------------
// Initialises an Obstacle, setting its position, size and health.
func CreateNewObstacle(xPos, yPos float32, colour color.RGBA) *Obstacle {
	obstacle := Obstacle{xPos: xPos, yPos: yPos, size: obstacleSize, colour: colour}

	obstacle.id = currentObstacleId
	currentObstacleId += 1
	if currentObstacleId > maxId {
		currentObstacleId = 0
	}

	obstacle.initHealth()
	obstacle.PerformMove()

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
func (obstacle *Obstacle) PerformMove() {
	chance := rand.Int()%2 == 0
	dir := rand.IntN(4000)

	if chance {
		if dir < 1000 {
			obstacle.xPos -= 1.0
		} else if dir > 3000 {
			obstacle.xPos += 1.0
		}

		if dir > 3000 {
			obstacle.yPos += 1.0
		} else if dir > 1000 {
			obstacle.yPos -= 1.0
		}
	}

	obstacle.nextMove = rand.IntN(60) + 10
}

// ----------------------------------------------------------------------------
func (obstacle *Obstacle) Update() {
	obstacle.healthTimer += 1

	if obstacle.healthTimer > 200 {
		obstacle.healthTimer = 0

		if obstacle.health < obstacle.maxHealth {
			obstacle.TakeHit(-1)
		}
	}

	obstacle.nextMove -= 1

	if obstacle.nextMove <= 0 {
		obstacle.PerformMove()
	}
}

// ----------------------------------------------------------------------------
// Renders the Bouncer on to the provided screen
func (obstacle *Obstacle) Draw(screen *ebiten.Image) {
	obstacle.size = obstacleSize * (float32(obstacle.health*100/obstacleInitialHealth) / 100)
	vector.DrawFilledRect(screen, obstacle.xPos, obstacle.yPos, obstacle.size, obstacle.size, obstacle.colour, true)
}
