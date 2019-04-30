package read

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRound(ctx *gin.Context) {
	roundId := ctx.Param("roundId")
	if roundId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	round := persistence.FetchRound(roundId)
	ctx.JSON(200, &round)
}
