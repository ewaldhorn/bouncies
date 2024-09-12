package main

import (
	"image"
	"image/color"

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
func drawArc(screen *ebiten.Image, xPos, yPos, radius, startAngle, endAngle float32) {
	var path vector.Path

	path.MoveTo(xPos, yPos)
	path.Arc(xPos, yPos, radius, startAngle, endAngle, vector.Clockwise)
	path.Close()
	var vs []ebiten.Vertex
	var is []uint16
	op1 := &vector.StrokeOptions{}
	op1.Width = 5
	op1.LineJoin = vector.LineJoinRound
	vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, op1)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xff / float32(0xff)
		vs[i].ColorG = 0xff / float32(0xff)
		vs[i].ColorB = 0xff / float32(0xff)
		vs[i].ColorA = 1
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.NonZero
	screen.DrawTriangles(vs, is, whiteSubImage, op)
}
