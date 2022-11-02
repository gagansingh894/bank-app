package util

import (
	"database/sql"
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

func RandomSQLNullString() sql.NullString {
	return sql.NullString{
		String: randomString(rand.Intn(10)),
		Valid:  true,
	}
}

func RandomSQLNullInt32() sql.NullInt32 {
	return sql.NullInt32{
		Int32: rand.Int31(),
		Valid: true,
	}
}

func RandomSQLNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Int64: rand.Int63(),
		Valid: true,
	}
}

func RandomSQLNullTime(time time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  time,
		Valid: true,
	}
}
