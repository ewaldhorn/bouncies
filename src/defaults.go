package main

import (
	"image/color"
)

const DEFAULT_HOMEBASE_HEALTH int = 1000
const DEFAULT_BASE_COUNT int = 2
const DEFAULT_BASE_OFFSET_BUFFER float32 = 5.0
const DEFAULT_TICKS_PER_BOUNCER_RESPAWN int = 180
const DEFAULT_TICKS_PER_SHIELD_REGEN int = 15
const DEFAULT_MAX_BOUNCERS int = 6

// ----------------------------------------------------------------------------
// Game colours
var COLOUR_RED = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var COLOUR_GREEN = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var COLOUR_BLUE = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var COLOUR_DARK_GRAY = color.RGBA{R: 64, G: 64, B: 64, A: 255}
