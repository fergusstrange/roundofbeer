package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidRoundPathParam(t *testing.T) {
	roundId := ValidRoundPathParam(NewTestContext("12345", "abcdef"))

	assert.Equal(t, "12345", roundId)
}

func TestValidRoundPathParam_InvalidParam(t *testing.T) {
	testContext := NewTestContext("", "")

	roundId := ValidRoundPathParam(testContext)

	assert.Equal(t, "", roundId)
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
	RoundId     string
	XRoundToken string
}

func NewTestContext(roundId, xRoundToken string) *TestContext {
	return &TestContext{
		RoundId:     roundId,
		XRoundToken: xRoundToken,
	}
}

func (ctx *TestContext) AbortWithStatus(code int) {
	ctx.AbortStatus = code
}

func (ctx *TestContext) Param(name string) string {
	return ctx.RoundId
}

func (ctx *TestContext) GetHeader(name string) string {
	return ctx.XRoundToken
}
