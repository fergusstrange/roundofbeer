package read

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

func Round(ctx *gin.Context) {
	roundToken, err := jwt.NewHelper().Decode(validation.ValidRoundHeader(ctx))
	if err != nil {
		ctx.AbortWithStatus(400)
	} else if fetchedRound := persistence.FetchRound(roundToken.RoundUrl); fetchedRound != nil {
		ctx.JSON(200, round.TransformRound(fetchedRound))
	} else {
		ctx.AbortWithStatus(400)
	}
}
