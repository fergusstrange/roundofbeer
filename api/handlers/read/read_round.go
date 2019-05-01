package read

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
)

func Round(ctx *gin.Context) {
	roundToken, err := jwt.NewHelper().Decode(validation.ValidRoundHeader(ctx))
	if err != nil {
		ctx.AbortWithStatus(400)
	} else if round := persistence.FetchRound(roundToken.RoundUrl); round != nil {
		ctx.JSON(200, validation.Transform(round))
	} else {
		ctx.AbortWithStatus(400)
	}
}
