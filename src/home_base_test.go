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
func TestHomeBaseAbsorbShield(t *testing.T) {
	damageToTake := 750
	shieldToAbsorb := 100
	absorbedShield := shieldToAbsorb / 2

	base := createKnownHomeBase()
	base.TakeDamage(damageToTake)

	if base.health != DEFAULT_HOMEBASE_HEALTH-damageToTake {
		t.Errorf("Expected health to be %d, but got %d", DEFAULT_HOMEBASE_HEALTH-damageToTake, base.health)
	}

	base.AbsorbShield(shieldToAbsorb) // remember absorb only takes half
	if base.health != DEFAULT_HOMEBASE_HEALTH-damageToTake+absorbedShield {
		t.Errorf("Expected health to be %d, but got %d", DEFAULT_HOMEBASE_HEALTH-damageToTake+absorbedShield, base.health)
	}

	base.AbsorbShield(5000)
	if base.health != DEFAULT_HOMEBASE_HEALTH {
		t.Errorf("Expected health to be %d, but got %d", DEFAULT_HOMEBASE_HEALTH, base.health)
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

// ----------------------------------------------------------------------------
func TestHomeBaseCreatePlayerBase(t *testing.T) {
	base := createPlayerHomeBase()

	if base.side != PLAYER_SIDE {
		t.Errorf("Expected base to be the player base, it wasn't")
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseCreateEnemyBase(t *testing.T) {
	base := createEnemyHomeBase()

	if base.side != ENEMY_SIDE {
		t.Errorf("Expected base to be the enemy base, it wasn't")
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseAdjustEnemyAttackAngle(t *testing.T) {
	base := createEnemyHomeBase()
	startAngle := base.attackAngle

	base.AdjustEnemyAttackAngle(20)
	if base.attackAngle != startAngle {
		t.Errorf("Expected attack angle to not change from %f to %f", startAngle, base.attackAngle)
	}

	base.AdjustEnemyAttackAngle(19)
	if base.attackAngle == startAngle {
		t.Errorf("Expected attack angle to change from %f", startAngle)
	}

	base.AdjustEnemyAttackAngle(80)
	if base.attackAngle == startAngle {
		t.Errorf("Expected attack angle to not change from %f to %f", startAngle, base.attackAngle)
	}

	startAngle = base.attackAngle
	base.AdjustEnemyAttackAngle(100)
	if base.attackAngle == startAngle {
		t.Errorf("Expected attack angle to change from %f (%f)", startAngle, base.attackAngle)
	}

	base.AdjustAttackAngle(-200)
	if base.attackAngle != ATTACK_ANGLE_MIN {
		t.Errorf("Expected attack angle to be %f, got %f", ATTACK_ANGLE_MIN, base.attackAngle)
	}

	base.AdjustAttackAngle(200)
	if base.attackAngle != ATTACK_ANGLE_MAX {
		t.Errorf("Expected attack angle to be %f, got %f", ATTACK_ANGLE_MAX, base.attackAngle)
	}
}

// ----------------------------------------------------------------------------
func TestHomeBaseFireBouncer(t *testing.T) {
	base := createDefaultHomeBase()

	didFire, bouncer := base.FireBouncer()

	if didFire {
		t.Errorf("Did not expect a bouncer to be fired!")
	}

	if bouncer != nil {
		t.Errorf("Did not expect a bouncer to be returned!")
	}

	base.bouncersAvailable = 1
	didFire, bouncer = base.FireBouncer()

	if !didFire {
		t.Errorf("Expect a bouncer to be fired!")
	}

	if bouncer == nil {
		t.Errorf("Expect a bouncer to be returned!")
	}

	if base.bouncersAvailable != 0 {
		t.Errorf("Expected no bouncers to be available, found %d", base.bouncersAvailable)
	}
}
