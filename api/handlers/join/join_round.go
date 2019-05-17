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

func NewJoinRoundHandler(serviceHandler func(roundId string, request *RoundRequest) (*round.WithToken, int)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		joinRoundRequest := new(RoundRequest)
		errors.LogError(ctx.BindJSON(joinRoundRequest))
		roundId := validation.ValidRoundPathParam(ctx)
		joinedRound, status := serviceHandler(roundId, joinRoundRequest)
		ctx.JSON(status, joinedRound)
	}
}

func Round(roundId string, request *RoundRequest) (*round.WithToken, int) {
	if fetchedRound := persistence.FetchRound(roundId); fetchedRound != nil &&
		nameNotAlreadyExists(request, fetchedRound.Participants) {
		token := round.EncodeRoundToken(fetchedRound.Url)
		return &round.WithToken{
			Token: &token,
			RoundUrl: &fetchedRound.Url,
		}, 200
	}
	return nil, 400
}

func nameNotAlreadyExists(request *RoundRequest, participants []persistence.Participant) bool {
	for _, participant := range participants {
		if strings.EqualFold(participant.Name, request.Name) {
			return false
		}
	}
	return true
}
