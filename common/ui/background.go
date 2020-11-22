package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juan-medina/goecs"
)

// BackgroundType is the ComponentType of Background
var BackgroundType = goecs.NewComponentType()

// Type will return Button goecs.ComponentType
func (b Background) Type() goecs.ComponentType {
	return BackgroundType
}

// NewBackground is
func NewBackground(x int, y int, image *ebiten.Image) *Background {
	return &Background{
		X:     x,
		Y:     y,
		image: image,
	}
}

// Background represents background.
type Background struct {
	X     int
	Y     int
	image *ebiten.Image
}

// Update ...
func (b *Background) Update() {
}

// setupGeom is reset, scale, translate geoM
func (b *Background) setupGeom(op *ebiten.DrawImageOptions, x int, y int) {
	op.GeoM.Reset()
	// op.GeoM.Scale(1, 2)
	op.GeoM.Translate(float64(x), float64(y))
}

// Draw renders content
func (b *Background) Draw(dst *ebiten.Image, op *ebiten.DrawImageOptions) {
	b.setupGeom(op, b.X, b.Y)
	dst.DrawImage(b.image, op)
}

// GetImage ...
func (b *Background) GetImage() *ebiten.Image {
	if b.image == nil {
		return nil
	}
	return b.image
}
