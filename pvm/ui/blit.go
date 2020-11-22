package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// CopyPixel copy
func CopyPixel(image *ebiten.Image, pXsrc int, pYsrc int, pXdst int, pYdst int) {
	c := image.At(pXsrc, pYsrc)
	image.Set(pXdst, pYdst, c)
}

// Floodfill fill
func Floodfill(image *ebiten.Image, x int, y int, c color.Color) {
	if x < image.Bounds().Min.X || image.Bounds().Max.X <= x || y < image.Bounds().Min.Y || image.Bounds().Max.Y <= y {
		return
	}
	targetColor := image.At(x, y)
	if targetColor == c {
		return
	}
	floodfillInternal(image, x, y, c, targetColor)
}

func floodfillInternal(image *ebiten.Image, x int, y int, c color.Color, targetColor color.Color) {
	if x < image.Bounds().Min.X || image.Bounds().Max.X <= x || y < image.Bounds().Min.Y || image.Bounds().Max.Y <= y {
		return
	}
	currentColor := image.At(x, y)
	if currentColor == targetColor {
		image.Set(x, y, c)
		floodfillInternal(image, x-1, y, c, targetColor)
		floodfillInternal(image, x+1, y, c, targetColor)
		floodfillInternal(image, x, y-1, c, targetColor)
		floodfillInternal(image, x, y+1, c, targetColor)
	}
}

// DrawRectangle draw
func DrawRectangle(screen *ebiten.Image, x1 int, y1 int, x2 int, y2 int, c color.Color) {
	ebitenutil.DrawLine(screen, float64(x1), float64(y1), float64(x2), float64(y1), c)
	ebitenutil.DrawLine(screen, float64(x1), float64(y2), float64(x2), float64(y2), c)
	ebitenutil.DrawLine(screen, float64(x1), float64(y1), float64(x1), float64(y2), c)
	ebitenutil.DrawLine(screen, float64(x2), float64(y1), float64(x2), float64(y2), c)
}
