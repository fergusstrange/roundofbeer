package pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	lol := String("lol")
	expected := "lol"

	assert.Equal(t, &expected, lol)
}
