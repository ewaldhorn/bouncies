package main

import "testing"

// ----------------------------------------------------------------------------
func TestObstacle_Init(t *testing.T) {
	currentObstacleId = maxId
	var xPos float32 = 123.0
	var yPos float32 = 456.0

	obstacle := CreateNewObstacle(xPos, yPos, COLOUR_GREEN)

	if currentObstacleId != 0 {
		t.Errorf("Expected the currentObstacleId to be %d, it was %d", 0, currentId)
	}

	if !almostEqual(float64(obstacle.xPos), float64(xPos)) || !almostEqual(float64(obstacle.yPos), float64(yPos)) {
		t.Errorf("Expected the obstacle XY to be (%f,%f) instead of (%f,%f)", xPos, yPos, obstacle.xPos, obstacle.yPos)
	}

	if obstacle.health != obstacleInitialHealth || obstacle.maxHealth != obstacle.health {
		t.Errorf("Expected obstacle health and maxHealth to be %d, not %d and %d", obstacleInitialHealth, obstacle.health, obstacle.maxHealth)
	}
}

// ----------------------------------------------------------------------------
func TestObstacle_TakeHit(t *testing.T) {
	obstacle := CreateNewObstacle(0.0, 0.0, COLOUR_WHITE)

	if obstacle.health != obstacleInitialHealth {
		t.Errorf("Expected health of %d, got %d", obstacleInitialHealth, obstacle.health)
	}

	obstacle.TakeHit(-100)
	if obstacle.health != obstacleInitialHealth {
		t.Errorf("Expected health of %d, got %d", obstacleInitialHealth, obstacle.health)
	}

	obstacle.TakeHit(11)
	if obstacle.health != obstacleInitialHealth-11 {
		t.Errorf("Expected health of %d, got %d", obstacleInitialHealth-11, obstacle.health)
	}

	obstacle.TakeHit(711)
	if obstacle.health != 0 {
		t.Errorf("Expected health of %d, got %d", 0, obstacle.health)
	}
}

// ----------------------------------------------------------------------------
func TestObstacle_Update(t *testing.T) {
	obstacle := CreateNewObstacle(10.0, 10.0, COLOUR_DARK_BLUE)

	preTimer := obstacle.healthTimer

	obstacle.Update()
	if obstacle.healthTimer <= preTimer {
		t.Errorf("Expected the health timer to increase from %d, got %d", preTimer, obstacle.healthTimer)
	}

	obstacle.healthTimer = 200
	obstacle.nextMove = 1
	obstacle.health -= 1

	obstacle.Update()
	if obstacle.healthTimer != 0 {
		t.Errorf("Expected the health timer to be %d, got %d", 0, obstacle.healthTimer)
	}

	if obstacle.nextMove < 50 {
		t.Errorf("Expected the next move to be %d+ away, got %d", 50, obstacle.nextMove)
	}

	if obstacle.health != obstacleInitialHealth {
		t.Errorf("Expected health to be %d, got %d", obstacleInitialHealth, obstacle.health)
	}

}
