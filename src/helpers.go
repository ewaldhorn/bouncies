package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ----------------------------------------------------------------------------
// See https://ebitengine.org/en/examples/vector.html for more information
var (
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

// ----------------------------------------------------------------------------
func init() {
	whiteImage.Fill(color.White)
}

// ----------------------------------------------------------------------------
// Draws an arc on the given screen element
func drawArc(screen *ebiten.Image, centerX, centerY, radius, width, startAngle, endAngle float32, clr color.Color) {

	if radius <= 0 || width <= 0 {
		log.Println("Invalid radius or width")
		return
	}

	var path vector.Path

	path.MoveTo(centerX, centerY)
	path.Arc(centerX, centerY, radius, startAngle, endAngle, vector.Clockwise)
	path.Close()

	strokeOptions := &vector.StrokeOptions{}
	strokeOptions.Width = width
	strokeOptions.LineJoin = vector.LineJoinRound
	vertices, indices := path.AppendVerticesAndIndicesForStroke(nil, nil, strokeOptions)

	r, g, b, a := clr.RGBA()
	for position := range vertices {
		vertices[position].SrcX = 1
		vertices[position].SrcY = 1
		vertices[position].ColorR = float32(r) / 0xffff
		vertices[position].ColorG = float32(g) / 0xffff
		vertices[position].ColorB = float32(b) / 0xffff
		vertices[position].ColorA = float32(a) / 0xffff
	}

	drawOptions := &ebiten.DrawTrianglesOptions{}
	drawOptions.FillRule = ebiten.NonZero
	screen.DrawTriangles(vertices, indices, whiteSubImage, drawOptions)
}

// ----------------------------------------------------------------------------
// Draws a filled arc
func drawFilledArc(screen *ebiten.Image, centerX, centerY, radius, startAngle, endAngle float32, clr color.Color) {
	var path vector.Path
	path.Arc(centerX, centerY, radius, startAngle, endAngle, vector.Clockwise)
	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	r, g, b, a := clr.RGBA()
	for position := range vertices {
		vertices[position].SrcX = 1
		vertices[position].SrcY = 1
		vertices[position].ColorR = float32(r) / 0xffff
		vertices[position].ColorG = float32(g) / 0xffff
		vertices[position].ColorB = float32(b) / 0xffff
		vertices[position].ColorA = float32(a) / 0xffff
	}

	drawOptions := &ebiten.DrawTrianglesOptions{}
	drawOptions.ColorScaleMode = ebiten.ColorScaleModePremultipliedAlpha
	drawOptions.AntiAlias = true
	screen.DrawTriangles(vertices, indices, whiteSubImage, drawOptions)
}

// ----------------------------------------------------------------------------
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
