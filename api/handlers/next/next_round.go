package next

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Participant string `json:"participant"`
}

type ServiceContext struct {
	persistence persistence.Persistence
}

func Handler(serviceHandler func(roundTokenHeader string) (*round.Round, int)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		roundHeader := validation.ValidRoundHeader(ctx)
		body, status := serviceHandler(roundHeader)
		ctx.JSON(status, body)
	}
}

func NewServiceContext(db persistence.Persistence) ServiceContext {
	return ServiceContext{
		persistence: db,
	}
}

func (sc ServiceContext) ServiceHandler(roundTokenHeader string) (*round.Round, int) {
	roundToken, err := jwt.NewHelper().Decode(roundTokenHeader)
	if err != nil {
		return nil, 400
	} else if currentRound := sc.persistence.FetchRound(roundToken.RoundUrl); currentRound != nil {
		updatedRound := round.UpdatedRoundWithNextCandidate(currentRound)
		sc.persistence.UpdateParticipantsAndCurrentCandidate(updatedRound)
		return round.TransformRound(updatedRound), 200
	} else {
		return nil, 400
	}
}
