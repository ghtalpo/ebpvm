package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/juan-medina/goecs"
)

// TextType is the ComponentType of Text
var TextType = goecs.NewComponentType()

// Type will return Button goecs.ComponentType
func (t Text) Type() goecs.ComponentType {
	return TextType
}

// NewText ...
func NewText(x int, y int, text string, c color.Color, isCenter bool) *Text {
	return &Text{
		X:        x,
		Y:        y,
		Text:     text,
		Color:    c,
		IsCenter: isCenter,
	}
}

// Text represents text ui.
type Text struct {
	X        int
	Y        int
	Text     string
	Color    color.Color
	IsCenter bool
}

// Update ...
func (t *Text) Update() {
}

// Draw renders content
func (t *Text) Draw(dst *ebiten.Image) {
	xSub := 0
	if t.IsCenter {
		rect := text.BoundString(GetUIFont(), t.Text)
		xSub = rect.Dx() / 2
	}
	text.Draw(dst, t.Text, GetUIFont(), t.X-xSub, t.Y+GetTextYOffset(), t.Color)
}

// GetTextYOffset ...
func GetTextYOffset() int {
	return lineHeight - (lineHeight-uiFontMHeight)/2
}
