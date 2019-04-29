package read

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
)

func GetRound(ctx *gin.Context) {
	roundId := ctx.Param(":roundId")
	round := persistence.FetchRound(roundId)
	ctx.JSON(200, round)
}
