package main

import "math"

const (
	// Player sides
	PLAYER_SIDE int = 0
	ENEMY_SIDE  int = 1
)

const (
	// Base constants
	DEFAULT_HOMEBASE_HEALTH           int     = 1000
	DEFAULT_BASE_COUNT                int     = 2
	DEFAULT_BASE_OFFSET_BUFFER        float32 = 2.0
	DEFAULT_MAX_BOUNCERS              int     = 8
	DEFAULT_FIRE_DELAY                int     = 80
	DEFAULT_TICKS_PER_BOUNCER_RESPAWN int     = 70
	DEFAULT_TICKS_PER_SHIELD_REGEN    int     = 40
)

const (
	// Math constants
	RADIAN           float32 = (math.Pi / 180)
	ATTACK_ANGLE_MIN         = -120.0
	ATTACK_ANGLE_MAX         = 28.0
)
