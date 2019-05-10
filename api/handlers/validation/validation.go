package validation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidRoundPathParam(ctx *gin.Context) string {
	roundId := ctx.Param("roundId")
	if roundId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundId
}

func ValidRoundHeader(ctx *gin.Context) string {
	roundToken := ctx.GetHeader("x-round-token")
	if roundToken == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundToken
}