package mathutil

import (
	"math"
)

// LoNibble returns lo nibble
func LoNibble(v int) int {
	return v & 0x0f
}

// HiNibble returns hi nibble
func HiNibble(v int) int {
	return (v >> 4) & 0x0f
}

// MakeByte returns a new byte (hi + lo)
func MakeByte(hi int, lo int) byte {
	return byte((LoNibble(hi) << 4) | LoNibble(lo))
}

// LoByte returns lo byte
func LoByte(v int) byte {
	return byte(v & 0xff)
}

// GetLittleEndianBytesFromWord ...
func GetLittleEndianBytesFromWord(v int) (b0 byte, b1 byte) {
	return byte(v & 0xff), byte((v >> 8) & 0xff)
}

// MakeWordFromLittleEndianBytes ...
func MakeWordFromLittleEndianBytes(b0 byte, b1 byte) int {
	return int(b0) | (int(b1) << 8)
}

// GetLittleEndianBytesFromDWord ...
func GetLittleEndianBytesFromDWord(v int) (b0 byte, b1 byte, b2 byte, b3 byte) {
	return byte(v & 0xff), byte((v >> 8) & 0xff), byte((v >> 16) & 0xff), byte((v >> 24) & 0xff)
}

// MakeDWordFromLittleEndianBytes ...
func MakeDWordFromLittleEndianBytes(b0 byte, b1 byte, b2 byte, b3 byte) int {
	return int(b0) | (int(b1) << 8) | (int(b2) << 16) | (int(b3) << 24)
}

// XClamp returns a value bounded with mi and ma
func XClamp(v int, mi int, ma int) int {
	if v < mi {
		v = mi
	}
	if ma < v {
		v = ma
	}
	return v
}

// XFloor returns a value lower bounded by l
func XFloor(v int, l int) int {
	return int(Max2(v, l))
}

// XCeil returns a value upper bounded by l
func XCeil(v int, l int) int {
	return int(Min2(v, l))
}

// Min2 returns minimum between 2 values
func Min2(v0 int, v1 int) int {
	if v0 < v1 {
		return v0
	}
	return v1
}

// Min3 returns minimum between 3 values
func Min3(v0 int, v1 int, v2 int) int {
	return int(Min2(Min2(v0, v1), v2))
}

// Min4 returns minimum between 4 values
func Min4(v0 int, v1 int, v2 int, v3 int) int {
	return int(Min2(Min2(v0, v1), Min2(v2, v3)))
}

// Min5 returns minimum between 5 values
func Min5(v0 int, v1 int, v2 int, v3 int, v4 int) int {
	return int(Min2(Min4(v0, v1, v2, v3), v4))
}

// Max2 returns maximum between 2 values
func Max2(v0 int, v1 int) int {
	if v0 < v1 {
		return v1
	}
	return v0
}

// Max3 returns maximum between 3 values
func Max3(v0 int, v1 int, v2 int) int {
	return int(Max2(Max2(v0, v1), v2))
}

// Max4 returns maximum between 4 values
func Max4(v0 int, v1 int, v2 int, v3 int) int {
	return int(Max2(Max2(v0, v1), Max2(v2, v3)))
}

// XAbs returns |v|
func XAbs(v int) int {
	return int(math.Abs(float64(v)))
}

func ternaryOpByte(b bool, trueV byte, falseV byte) byte {
	if b {
		return trueV
	}
	return falseV
}

func ternaryOpInt(b bool, trueV int, falseV int) int {
	if b {
		return trueV
	}
	return falseV
}

// TernaryOpInt is a external wrapper for ternaryOpInt
func TernaryOpInt(b bool, trueV int, falseV int) int {
	return ternaryOpInt(b, trueV, falseV)
}

func ternaryOpString(b bool, trueV string, falseV string) string {
	if b {
		return trueV
	}
	return falseV
}

// TernaryOpString is a external wrapper for ternaryOpString
func TernaryOpString(b bool, trueV string, falseV string) string {
	return ternaryOpString(b, trueV, falseV)
}
