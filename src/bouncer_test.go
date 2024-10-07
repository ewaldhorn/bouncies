package main

import "testing"

// ----------------------------------------------------------------------------
func TestBouncer_Init(t *testing.T) {
	currentId = maxId
	base := createDefaultHomeBase()
	bouncer := Bouncer{}
	bouncer.Init(*base)

	if currentId != 0 {
		t.Errorf("Expected the currentId to be %d, it was %d", 0, currentId)
	}
}

// ----------------------------------------------------------------------------
func TestBouncer_TakeHit(t *testing.T) {
	bouncer := Bouncer{}
	bouncer.health = 10
	bouncer.maxHealth = initialHealth

	bouncer.TakeHit(11)
	if bouncer.health != 0 {
		t.Errorf("1. Bouncer health should be 0, not %d", bouncer.health)
	}

	bouncer.health = 11
	bouncer.TakeHit(11)
	if bouncer.health != 0 {
		t.Errorf("2. Bouncer health should be 0, not %d", bouncer.health)
	}

	bouncer.health = 10
	bouncer.TakeHit(9)
	if bouncer.health != 1 {
		t.Errorf("3. Bouncer health should be 1, not %d", bouncer.health)
	}

	bouncer.health = initialHealth
	bouncer.TakeHit(-11)
	if bouncer.health != initialHealth {
		t.Errorf("4. Bouncer health should be %d, not %d", initialHealth, bouncer.health)
	}

}
