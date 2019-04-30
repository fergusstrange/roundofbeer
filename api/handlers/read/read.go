package read

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
)

func GetRound(ctx *gin.Context) {
	roundId := validation.ValidRoundId(ctx)
	round := persistence.FetchRound(roundId)
	ctx.JSON(200, &round)
}
