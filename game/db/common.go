package db

import (
	"image/color"
)

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

// Color is temporal wrapper
func Color(r int, g int, b int) color.Color {
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}
