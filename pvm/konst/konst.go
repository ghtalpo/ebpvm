package konst

// CheatEnabled is
const CheatEnabled = true

const (
	// Key0 is
	Key0 = 0
	// Key1 is
	Key1 = 1
	// Key2 is
	Key2 = 2
	// Key3 is
	Key3 = 3
	// Key4 is
	Key4 = 4
	// Key5 is
	Key5 = 5
	// Key6 is
	Key6 = 6
	// Key7 is
	Key7 = 7
	// Key8 is
	Key8 = 8
	// Key9 is
	Key9 = 9
	// KeyEnter is
	KeyEnter = 13
	// KeyU is
	KeyU = -1 // gara for now
	// KeyI is
	KeyI = -1 // gara for now
	// KeyJ is
	KeyJ = -1 // gara for now
	// KeyO is
	KeyO = -1 // gara for now
	// KeyK is
	KeyK = -1 // gara for now
	// KeyL is
	KeyL = -1 // gara for now
	// KeyM is
	KeyM = -1 // gara for now
	// KeySlash is
	KeySlash = -1 // gara for now
	// KeyAsterisk is
	KeyAsterisk = -1 // gara for now
	// KeyF10 is
	KeyF10 = -1 // gara for now
)

// NumberInvalid is
const NumberInvalid = -1

// FrameRate for converting tick to seconds
const FrameRate = 30

// OK is
const OK = 1

// FakeType is
type FakeType int

const (
	// FakeTypeNull is
	FakeTypeNull = iota
	// FakeTypeInt is
	FakeTypeInt
	// FakeTypeWar is
	FakeTypeWar
	// FakeTypeChr is
	FakeTypeChr
)

// GetFakeType is
func GetFakeType(n int) FakeType {
	switch n {
	case int(FakeTypeNull):
		return FakeTypeNull
	case int(FakeTypeInt):
		return FakeTypeInt
	case int(FakeTypeWar):
		return FakeTypeWar
	case int(FakeTypeChr):
		return FakeTypeChr
	default:
		panic(n)
	}
}
