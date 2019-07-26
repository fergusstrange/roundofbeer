package random

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestAlphaNumeric(t *testing.T) {
	alphaNumeric := AlphaNumeric(12)

	assert.MatchRegex(t, alphaNumeric, "^[abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]{12}$")
}
