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
