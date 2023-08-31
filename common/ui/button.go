package ui

import (
	"image"

	"github.com/juan-medina/goecs"
)

// ButtonType is the ComponentType of Button
var ButtonType = goecs.NewComponentType()

// Type will return Button goecs.ComponentType
func (b Button) Type() goecs.ComponentType {
	return ButtonType
}

// NewButton ...
func NewButton(rect image.Rectangle, text string, onPressed func(b *Button)) *Button {
	return &Button{
		Rect:      rect,
		Text:      text,
		onPressed: onPressed,
		enabled: true,
	}
}
