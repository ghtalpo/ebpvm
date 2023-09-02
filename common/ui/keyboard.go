package ui

import (
	"image"
	"image/color"
	_ "image/png" //
	"log"
	"strings"
	"os"

	"github.com/ghtalpo/egb/common/ui/keyboard"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var keyboardImage *ebiten.Image
var capsImage *ebiten.Image
// var rect image.Rectangle
var validKeys []ebiten.Key = []ebiten.Key{
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
	ebiten.KeyA,
	ebiten.KeyB,
	ebiten.KeyC,
	ebiten.KeyD,
	ebiten.KeyE,
	ebiten.KeyF,
	ebiten.KeyG,
	ebiten.KeyH,
	ebiten.KeyI,
	ebiten.KeyJ,
	ebiten.KeyK,
	ebiten.KeyL,
	ebiten.KeyM,
	ebiten.KeyN,
	ebiten.KeyO,
	ebiten.KeyP,
	ebiten.KeyQ,
	ebiten.KeyR,
	ebiten.KeyS,
	ebiten.KeyT,
	ebiten.KeyU,
	ebiten.KeyV,
	ebiten.KeyW,
	ebiten.KeyX,
	ebiten.KeyY,
	ebiten.KeyZ,
	ebiten.KeyApostrophe,
	ebiten.KeyBackslash,
	ebiten.KeyBackspace,
	ebiten.KeyComma,
	ebiten.KeyEnter,
	ebiten.KeyEqual,
	ebiten.KeyGraveAccent,
	ebiten.KeyMinus,
	ebiten.KeyPeriod,
	ebiten.KeyLeftBracket,
	ebiten.KeyRightBracket,
	ebiten.KeySemicolon,
	ebiten.KeySlash,
	ebiten.KeySpace,
	ebiten.KeyShift,
}

// LoadImage ...
func LoadImage(path string) *ebiten.Image {
	img, _ := loadImage(path)
	return img
}

// func
func loadImage(path string) (*ebiten.Image, bool) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage := ebiten.NewImageFromImage(img)
	return ebitenImage, true
}

func init() {
	keyboardImage, _ = loadImage("_resources/common/image/keyboard2.png")
	capsImage, _ = loadImage("_resources/common/image/caps.png")
}

// Keyboard ...
type Keyboard struct {
	TopLeft         image.Point
	feedBackPos     image.Point
	// pressed         []ebiten.Key
	toCaps          bool
	buffer          string
	result          string
	feedBackWidthB  int
	feedbackEnabled bool

	onPressed func(k *Keyboard)
}

// NewKeyboard is a constructor.
func NewKeyboard(topLeft image.Point) *Keyboard {
	k := Keyboard{TopLeft: topLeft, feedbackEnabled: false}
	k.toCaps = true
	k.buffer = ""
	k.result = ""
	return &k
}

// EnableFeedback enables
func (k *Keyboard) EnableFeedback(pos image.Point, widthB int) {
	k.feedbackEnabled = true
	k.feedBackPos = pos
	k.feedBackWidthB = widthB
}

// DisableFeedback disables
func (k *Keyboard) DisableFeedback() {
	k.feedbackEnabled = false
}

// GetBuffer gets temporal string.
func (k *Keyboard) GetBuffer() string {
	return k.buffer
}

// GetString gets composed string.
func (k *Keyboard) GetString() string {
	return k.result
}

// Clear clears composed string.
func (k *Keyboard) Clear() {
	k.result = ""
}

// SetOnPressed register callback.
func (k *Keyboard) SetOnPressed(f func(k *Keyboard)) {
	k.onPressed = f
}

// Draw render textures
func (k *Keyboard) Draw(dst *ebiten.Image) {
	var (
		offsetX = k.TopLeft.X
		offsetY = k.TopLeft.Y
	)

	colorBorder := color.Color(color.RGBA{0x7f, 0x7f, 0x7f, 0xff})
	width, height := keyboardImage.Bounds().Dx(), keyboardImage.Bounds().Dy()

	// draw pretty borders
	// horz
	vector.StrokeLine(dst, float32(offsetX-1), float32(offsetY-2), float32(offsetX-1+width+2), float32(offsetY-2+0), 1, colorBorder, false)
	vector.StrokeLine(dst, float32(offsetX-1), float32(offsetY-1), float32(offsetX-1+width+2), float32(offsetY-1+0), 1, color.Black, false)
	// vert
	vector.StrokeLine(dst, float32(offsetX-1), float32(offsetY-1), float32(offsetX-1), float32(offsetY-2+height+2), 1, colorBorder, false)
	vector.StrokeLine(dst, float32(offsetX), float32(offsetY-1), float32(offsetX), float32(offsetY-2+height+2), 1, color.Black, false)
	vector.StrokeLine(dst, float32(offsetX-2+width+3), float32(offsetY-1), float32(offsetX-2+width+3), float32(offsetY-2+height+2), 1, color.Black, false)
	vector.StrokeLine(dst, float32(offsetX-1+width+3), float32(offsetY-1), float32(offsetX-1+width+3), float32(offsetY-2+height+2), 1, colorBorder, false)
	// vector.StrokeLine(dst, float64(x), float64(80)*2, float64(x), float64(143)*2, colorBorder)

	// Draw the base (grayed) keyboard image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(offsetX), float64(offsetY))
	op.ColorScale.Scale(0.5, 0.5, 0.5, 1)
	dst.DrawImage(keyboardImage, op)

	// Draw caps lock status with shift
	if k.toCaps {
		r, ok := keyboard.KeyRect(ebiten.KeyShift)
		if ok {
			op.ColorScale.Reset()
			op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
			dst.DrawImage(capsImage, op)
		}
	}

	if k.feedbackEnabled {
		x := k.feedBackPos.X
		y := k.feedBackPos.Y
		text.Draw(dst, k.GetBuffer(), GetUIFont(), x, y, color.White)
		vector.StrokeLine(dst, float32(x), float32(y+2), float32(x+k.feedBackWidthB*8), float32(y+2), 1, color.White, false)
	}
}

// Update handles input. simple automata
func (k *Keyboard) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		var (
			offsetX = k.TopLeft.X
			offsetY = k.TopLeft.Y
		)
		x, y := ebiten.CursorPosition()
		point := image.Point{x - offsetX, y - offsetY}
		for _, key := range validKeys {
			r, _ := keyboard.KeyRect(key)
			if point.In(r) {
				switch key {
				case ebiten.KeyBackspace:
					if len(k.buffer) > 0 {
						k.buffer = k.buffer[:len(k.buffer)-1]
					}
				case ebiten.KeyShift:
					k.toCaps = !k.toCaps
				default:
					k.buffer += k.keyToString(key)
				}
			}
		}
	}
}

func (k *Keyboard) keyToString(key ebiten.Key) string {
	switch key {
	case ebiten.KeySpace:
		k.toCaps = true
		return " "
	case ebiten.KeyApostrophe:
		return "'"
	case ebiten.KeyBackslash:
		return "\\"
	case ebiten.KeyComma:
		return ","
	case ebiten.KeyEqual:
		return "="
	case ebiten.KeyGraveAccent:
		return "`"
	case ebiten.KeyMinus:
		return "-"
	case ebiten.KeyPeriod:
		return "."
	case ebiten.KeyLeftBracket:
		return "["
	case ebiten.KeyRightBracket:
		return "]"
	case ebiten.KeySemicolon:
		return ";"
	case ebiten.KeySlash:
		return "/"
	case ebiten.KeyEnter:
		k.result = k.buffer
		k.buffer = ""
		if k.onPressed != nil {
			k.onPressed(k)
		}
		return ""
	default:
		if k.toCaps {
			k.toCaps = false
			return strings.ToUpper(key.String())
		}
		return strings.ToLower(key.String())
	}
}
