package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return randomString(rand.Intn(10))
}

// generates random currency type
func RandomCurrency() string {
	currencies := []string{"USD", "GBP", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
