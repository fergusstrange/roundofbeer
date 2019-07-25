package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidRoundPathParam(t *testing.T) {
	roundID := ValidRoundPathParam(NewTestContext("12345", "abcdef"))

	assert.Equal(t, "12345", roundID)
}

func TestValidRoundPathParam_InvalidParam(t *testing.T) {
	testContext := NewTestContext("", "")

	roundID := ValidRoundPathParam(testContext)

	assert.Equal(t, "", roundID)
	assert.Equal(t, 400, testContext.AbortStatus)
}

func TestValidRoundHeader(t *testing.T) {
	roundToken := ValidRoundHeader(NewTestContext("12345", "abcdef"))

	assert.Equal(t, "abcdef", roundToken)
}

func TestValidRoundHeader_InvalidParam(t *testing.T) {
	testContext := NewTestContext("", "")

	roundToken := ValidRoundHeader(testContext)

	assert.Equal(t, "", roundToken)
	assert.Equal(t, 400, testContext.AbortStatus)
}

type TestContext struct {
	AbortStatus int
	RoundID     string
	XRoundToken string
}

func NewTestContext(roundID, xRoundToken string) *TestContext {
	return &TestContext{
		RoundID:     roundID,
		XRoundToken: xRoundToken,
	}
}

func (ctx *TestContext) AbortWithStatus(code int) {
	ctx.AbortStatus = code
}

func (ctx *TestContext) Param(name string) string {
	return ctx.RoundID
}

func (ctx *TestContext) GetHeader(name string) string {
	return ctx.XRoundToken
}
