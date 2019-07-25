package create

import (
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/random"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type Request struct {
	Participants []string `json:"participants"`
}

type ServiceContext struct {
	persistence persistence.Persistence
}

func Handler(serviceHandler func(request *Request) round.WithToken) func(*gin.Context) {
	return func(ctx *gin.Context) {
		createRoundRequest := new(Request)
		errors.LogError(ctx.BindJSON(createRoundRequest))
		ctx.JSON(200, serviceHandler(createRoundRequest))
	}
}

func NewServiceContext(db persistence.Persistence) ServiceContext {
	return ServiceContext{
		persistence: db,
	}
}

func (sc ServiceContext) ServiceHandler(request *Request) round.WithToken {
	url := random.AlphaNumeric(6)
	roundToPersist := round.UpdatedRoundWithNextCandidate(&persistence.Round{
		URL:          url,
		CreateDate:   time.Now(),
		UpdateDate:   time.Now(),
		Participants: transformParticipants(request),
	})
	sc.persistence.CreateRound(roundToPersist)
	return newRoundResponse(roundToPersist)
}

func transformParticipants(request *Request) []persistence.Participant {
	var participantList []persistence.Participant
	for _, participant := range request.Participants {
		participantList = append(participantList, persistence.Participant{
			UUID:       uuid.New().String(),
			Name:       participant,
			RoundCount: 0,
		})
	}
	return participantList
}

func newRoundResponse(persistedRound *persistence.Round) round.WithToken {
	encodedRoundToken := round.EncodeRoundToken(persistedRound.URL)
	return round.WithToken{
		Token:        &encodedRoundToken,
		RoundURL:     &persistedRound.URL,
		Participants: round.ParticipantsFromRound(persistedRound),
	}
}
