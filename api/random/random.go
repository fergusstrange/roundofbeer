package random

import (
	"math/rand"
	"time"
)

const alphaNumericChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func AlphaNumeric(length int) string {
	randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = alphaNumericChars[randomSeed.Intn(len(alphaNumericChars))]
	}
	return string(b)
}
