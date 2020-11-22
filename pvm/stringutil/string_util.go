package stringutil

import (
	"strconv"
	"strings"
)

// Itoa returns string representation of int
func Itoa(n int) string {
	return strconv.Itoa(n)
}

// Pad pads
func Pad(s string, n int) string {
	l := len(s)
	for i := 0; i < n-l; i++ {
		s = " " + s
	}
	return s
}

// Atoi returns
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// CleanString is
func CleanString(s string) string {
	i := strings.IndexByte(s, 0)
	if i != -1 {
		return s[:i]
	}
	return s
}
