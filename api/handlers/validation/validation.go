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
	roundID := ctx.Param("roundID")
	if roundID == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundID
}

func ValidRoundHeader(ctx Context) string {
	roundToken := ctx.GetHeader("x-round-token")
	if roundToken == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundToken
}
