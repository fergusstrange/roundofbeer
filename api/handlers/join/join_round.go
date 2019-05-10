package join

import (
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
	"strings"
)

type RoundRequest struct {
	Name string `json:"name" binding:"required"`
}

func Round(ctx *gin.Context) {
	roundId := validation.ValidRoundPathParam(ctx)
	joinRoundRequest := new(RoundRequest)
	errors.LogFatal(ctx.BindJSON(joinRoundRequest))
	fetchedRound := persistence.FetchRound(roundId)
	if fetchedRound != nil && nameNotAlreadyExists(joinRoundRequest, fetchedRound.Participants) {
		ctx.JSON(200, round.TransformRound(fetchedRound))
	} else {
		ctx.AbortWithStatus(400)
	}
}

func nameNotAlreadyExists(request *RoundRequest, participants []persistence.Participant) bool {
	for _, participant := range participants {
		if strings.EqualFold(participant.Name, request.Name) {
			return false
		}
	}
	return true
}
