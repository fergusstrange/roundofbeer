package validation

import (
	"net/http"
)

type Context interface {
	AbortWithStatus(code int)
	Param(name string) string
	GetHeader(name string) string
}

func ValidRoundPathParam(ctx Context) string {
	roundId := ctx.Param("roundId")
	if roundId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundId
}

func ValidRoundHeader(ctx Context) string {
	roundToken := ctx.GetHeader("x-round-token")
	if roundToken == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundToken
}
