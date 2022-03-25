package utils

import (
	"math/rand"
	"time"
)

func RandInt(min int, max int) uint64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return uint64(min + rand.Intn(max-min))
}
