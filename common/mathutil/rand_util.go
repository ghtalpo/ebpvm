package mathutil

import (
	"math/rand"
)

// RandSeed seeds
func RandSeed(seed int64) {
	rand.Seed(seed)
}

// RandByte returns [0, v)
func RandByte(v byte) byte {
	if v == 0 {
		return 0
	} else if v < 0 {
		panic("minus")
	}
	return byte(rand.Intn(int(v)))
}

// RandInt returns [0, v)
func RandInt(v int) int {
	if v == 0 {
		return 0
	} else if v < 0 {
		panic("minus")
	}
	return rand.Intn(v)
}

// RandInt2 returns [high, low)
func RandInt2(low int, high int) int {
	if !(low < high) {
		return int(low)
	}
	return rand.Intn(high-low) + low
}

// CheckPercent check if p percent of possiblity succeeded
func CheckPercent(p int) bool {
	return RandInt(100) < p
}

// PickStringEqually picks one from vector
func PickStringEqually(strVec []string) string {
	pick := RandInt(len(strVec))
	return strVec[pick]
}

// PickEqually picks one from vector
func PickEqually(strMap map[string]interface{}) string {
	keys := make([]string, len(strMap))
	i := 0
	for k := range strMap {
		keys[i] = k
		i++
	}

	pick := RandInt(len(strMap))
	key := keys[pick]
	return strMap[key].(string)
}
