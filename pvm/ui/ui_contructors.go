package ui

import (
	"image"

	"golang.org/x/image/font"
)

// NewCheckBox is
func NewCheckBox(x int, y int, text string) *CheckBox {
	return &CheckBox{
		X:    x,
		Y:    y,
		Text: text,
	}
}

// NewTextBox is
func NewTextBox(rect image.Rectangle) *TextBox {
	return &TextBox{
		Rect: rect,
	}
}

// GetUIFont is
func GetUIFont() font.Face {
	return uiFont
}
