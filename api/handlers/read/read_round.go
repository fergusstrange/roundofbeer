package read

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

type ServiceContext struct {
	persistence persistence.Persistence
}

func Handler(serviceHandler func(roundToken string) (*round.Round, int)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		roundTokenHeader := validation.ValidRoundHeader(ctx)
		body, status := serviceHandler(roundTokenHeader)
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
	} else if fetchedRound := sc.persistence.FetchRound(roundToken.RoundURL); fetchedRound != nil {
		return round.TransformRound(fetchedRound), 200
	} else {
		return nil, 400
	}
}
