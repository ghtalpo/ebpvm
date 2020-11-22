// Copyright 2013 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// copy from keyboard/keyrects.go

package numpad

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var keyboardKeyRects = map[ebiten.Key]image.Rectangle{}

func init() {
	keyboardKeyRects[ebiten.Key0] = image.Rect(0, 54, 16, 72)
	keyboardKeyRects[ebiten.Key1] = image.Rect(0, 36, 16, 54)
	keyboardKeyRects[ebiten.Key2] = image.Rect(16, 36, 32, 54)
	keyboardKeyRects[ebiten.Key3] = image.Rect(32, 36, 48, 54)
	keyboardKeyRects[ebiten.Key4] = image.Rect(0, 18, 16, 36)
	keyboardKeyRects[ebiten.Key5] = image.Rect(16, 18, 32, 36)
	keyboardKeyRects[ebiten.Key6] = image.Rect(32, 18, 48, 36)
	keyboardKeyRects[ebiten.Key7] = image.Rect(0, 0, 16, 18)
	keyboardKeyRects[ebiten.Key8] = image.Rect(16, 0, 32, 18)
	keyboardKeyRects[ebiten.Key9] = image.Rect(32, 0, 48, 18)
	keyboardKeyRects[ebiten.KeyBackspace] = image.Rect(48, 0, 64, 18)
	keyboardKeyRects[ebiten.KeyKPMultiply] = image.Rect(48, 18, 64, 36)
	keyboardKeyRects[ebiten.KeyKPDivide] = image.Rect(48, 36, 64, 54)
	keyboardKeyRects[ebiten.KeyEnter] = image.Rect(16, 54, 64, 72)
}

// KeyRect ...
func KeyRect(key ebiten.Key) (image.Rectangle, bool) {
	r, ok := keyboardKeyRects[key]
	return r, ok
}
