package ui

import (
	"image"
	"image/color"
	_ "image/png"

	"github.com/ghtalpo/egb/common/stringutil"
	"github.com/ghtalpo/egb/common/ui/numpad"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var numPadImage *ebiten.Image
var validNumPadKeys []ebiten.Key = []ebiten.Key{
	ebiten.Key0,
	ebiten.Key1,
	ebiten.Key2,
	ebiten.Key3,
	ebiten.Key4,
	ebiten.Key5,
	ebiten.Key6,
	ebiten.Key7,
	ebiten.Key8,
	ebiten.Key9,
	ebiten.KeyKPMultiply,
	ebiten.KeyKPDivide,
	ebiten.KeyBackspace,
	ebiten.KeyEnter,
}

func init() {
	numPadImage, _ = loadImage("_resources/common/image/numpad.png")
}

// NumPad ...
type NumPad struct {
	TopLeft         image.Point
	feedbackPos     image.Point
	pressed         []ebiten.Key
	buffer          string
	result          string
	min             int
	max             int
	allowEmpty      bool
	feedbackWidthB  int
	feedbackEnabled bool
	maxStringLength int
	feedbackColor   color.Color

	onPressed func(k *NumPad)
}

// NewNumPad is a constructor.
func NewNumPad(topLeft image.Point) *NumPad {
	k := NumPad{TopLeft: topLeft, min: -1, max: -1, allowEmpty: false, feedbackEnabled: false}
	k.buffer = ""
	k.result = ""
	k.feedbackColor = color.White
	return &k
}

// EnableFeedback enables
func (k *NumPad) EnableFeedback(pos image.Point, widthB int) {
	k.feedbackEnabled = true
	k.feedbackPos = pos
	k.feedbackWidthB = widthB
}

// DisableFeedback disables
func (k *NumPad) DisableFeedback() {
	k.feedbackEnabled = false
}

// SetFeedbackColor ...
func (k *NumPad) SetFeedbackColor(color color.Color) {
	k.feedbackColor = color
}

// GetBuffer gets temporal string.
func (k *NumPad) GetBuffer() string {
	return k.buffer
}

// GetString gets composed string.
func (k *NumPad) GetString() string {
	return k.result
}

// Clear clears composed string.
func (k *NumPad) Clear() {
	k.result = ""
}

// SetOnPressed register callback.
func (k *NumPad) SetOnPressed(f func(k *NumPad)) {
	k.onPressed = f
}

// SetMin sets lower bounds(minus means no check)
func (k *NumPad) SetMin(min int) {
	k.min = min
}

// SetMax sets upper bounds(minus means no check)
func (k *NumPad) SetMax(max int) {
	k.max = max
	k.maxStringLength = len(stringutil.Itoa(max))
}

// Draw render textures
func (k *NumPad) Draw(dst *ebiten.Image) {
	var (
		offsetX = k.TopLeft.X
		offsetY = k.TopLeft.Y
	)

	colorBorder := color.Color(color.RGBA{0x7f, 0x7f, 0x7f, 0xff})
	width, height := numPadImage.Size()
	// draw pretty borders
	// horz
	ebitenutil.DrawLine(dst, float64(offsetX-1), float64(offsetY-2), float64(offsetX-1+width+2), float64(offsetY-2+0), colorBorder)
	ebitenutil.DrawLine(dst, float64(offsetX-1), float64(offsetY-1), float64(offsetX-1+width+2), float64(offsetY-1+0), color.Black)
	// vert
	ebitenutil.DrawLine(dst, float64(offsetX-1), float64(offsetY-1), float64(offsetX-1), float64(offsetY-2+height+2), colorBorder)
	ebitenutil.DrawLine(dst, float64(offsetX), float64(offsetY-1), float64(offsetX), float64(offsetY-2+height+2), color.Black)
	ebitenutil.DrawLine(dst, float64(offsetX-2+width+3), float64(offsetY-1), float64(offsetX-2+width+3), float64(offsetY-2+height+2), color.Black)
	ebitenutil.DrawLine(dst, float64(offsetX-1+width+3), float64(offsetY-1), float64(offsetX-1+width+3), float64(offsetY-2+height+2), colorBorder)
	// ebitenutil.DrawLine(dst, float64(x), float64(80)*2, float64(x), float64(143)*2, colorBorder)

	// Draw the base (grayed) NumPad image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(offsetX), float64(offsetY))
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	dst.DrawImage(numPadImage, op)

	if k.feedbackEnabled {
		x := k.feedbackPos.X
		y := k.feedbackPos.Y
		text.Draw(dst, k.GetBuffer(), GetUIFont(), x, y, k.feedbackColor)
		ebitenutil.DrawLine(dst, float64(x), float64(y+2), float64(x+k.feedbackWidthB*8), float64(y+2), k.feedbackColor)
	}
}

// Update handles input. simple automata
func (k *NumPad) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var (
			offsetX = k.TopLeft.X
			offsetY = k.TopLeft.Y
		)
		x, y := ebiten.CursorPosition()
		point := image.Point{x - offsetX, y - offsetY}
		for _, key := range validNumPadKeys {
			r, _ := numpad.KeyRect(key)
			if point.In(r) {
				switch key {
				case ebiten.KeyBackspace:
					if len(k.buffer) > 0 {
						k.buffer = k.buffer[:len(k.buffer)-1]
					}
				case ebiten.KeyKPMultiply:
					k.buffer = stringutil.Itoa(k.max)
				case ebiten.KeyKPDivide:
					k.buffer = stringutil.Itoa(k.min)
				default:
					k.buffer += k.keyToString(key)
				}
			}
		}
	}
}

func (k *NumPad) keyToString(key ebiten.Key) string {
	switch key {
	case ebiten.KeyEnter:
		k.result = ""
		if k.buffer != "" {
			n := stringutil.Atoi(k.buffer)
			if k.min >= 0 {
				if n < k.min {
					k.buffer = ""
					return ""
				}
			}
			if k.max >= 0 {
				if k.max < n {
					k.buffer = ""
					return ""
				}
			}
			k.result = stringutil.Itoa(n)
		}
		k.buffer = ""
		if k.onPressed != nil {
			if k.result == "" && !k.allowEmpty {
			} else {
				k.onPressed(k)
			}
		}
		return ""
	default:
		if len(k.buffer) < k.maxStringLength {
			return key.String()
		}
		return ""
	}
}
