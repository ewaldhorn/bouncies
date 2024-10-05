package main

import (
	"testing"
)

const (
	KNOWN_X = 150
	KNOWN_Y = 125
)

// ----------------------------------------------------------------------------
func createKnownHomeBase() *HomeBase {
	return &HomeBase{centerX: 100, centerY: 101, radius: 50, health: DEFAULT_HOMEBASE_HEALTH}
}

// ----------------------------------------------------------------------------
func createDefaultHomeBase() *HomeBase {
	base := &HomeBase{}
	base.init(KNOWN_X, KNOWN_Y)
	return base
}

// ----------------------------------------------------------------------------
func TestHomeBaseCreation(t *testing.T) {
	base := createKnownHomeBase()

	if base.centerX != 100 {
		t.Errorf("Expected xPos to be 100, but got %f", base.centerX)
	}
	if base.centerY != 101 {
		t.Errorf("Expected yPos to be 100, but got %f", base.centerY)
	}
	if base.radius != 50 {
		t.Errorf("Expected radius to be 50, but got %f", base.radius)
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseInit(t *testing.T) {
	base := createDefaultHomeBase()

	if base.centerX != KNOWN_X || base.centerY != KNOWN_Y {
		t.Errorf("init failed base created with (%f,%f) as CenterX,CenterY, expected (%d,%d)", base.centerX, base.centerY, KNOWN_X, KNOWN_Y)
	}

	if base.health != base.maxHealth || base.maxHealth != DEFAULT_HOMEBASE_HEALTH {
		t.Errorf("init failed base created with health %d/%d, expected %d/%d", base.health, base.maxHealth, DEFAULT_HOMEBASE_HEALTH, DEFAULT_HOMEBASE_HEALTH)
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseTakeDamage(t *testing.T) {
	base := createKnownHomeBase()
	base.TakeDamage(50)
	if base.health != DEFAULT_HOMEBASE_HEALTH-50 {
		t.Errorf("Expected health to be %d, but got %d", DEFAULT_HOMEBASE_HEALTH-50, base.health)
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseTakeDamageOutsized(t *testing.T) {
	base := createKnownHomeBase()

	base.TakeDamage(10_000_000)

	if base.IsAlive() {
		t.Errorf("Expected base health to be 0, got %d", base.health)
	}

}

// ----------------------------------------------------------------------------
func TestHomeBaseIsAlive(t *testing.T) {
	base := createKnownHomeBase()
	base.health = 50

	if !base.IsAlive() {
		t.Errorf("Expected base to be alive, but it's not")
	}
	base.TakeDamage(50)
	if base.IsAlive() {
		t.Errorf("Expected base to be dead, but it's not")
	}
}
