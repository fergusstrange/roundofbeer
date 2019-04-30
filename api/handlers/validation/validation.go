package validation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidRoundId(ctx *gin.Context) string {
	roundId := ctx.Param("roundId")
	if roundId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundId
}
