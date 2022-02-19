package utils

import (
	"math/rand"
	"strings"
)

var pool = "aqwertyuioplkjhgfdsazxcvbm"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(pool)
	for i := 0; i < n; i++ {
		index := rand.Intn(k)
		sb.WriteByte(pool[index])
	}

	return sb.String()
}

func RandomInt() int64 {
	return rand.Int63()
}

func RandomFloat() float64 {
	num := rand.Float64()
	return num
}

func RandomUnit() uint32 {
	num := rand.Uint32()
	return num
}
