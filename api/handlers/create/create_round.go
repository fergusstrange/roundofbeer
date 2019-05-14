package create

import (
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/jwt"
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

type Response struct {
	Token *string      `json:"token"`
	Round *round.Round `json:"round"`
}

func NewRoundHandler(serviceHandler func(request *Request) Response) func(*gin.Context) {
	return func(ctx *gin.Context) {
		createRoundRequest := new(Request)
		errors.LogError(ctx.BindJSON(createRoundRequest))
		ctx.JSON(200, serviceHandler(createRoundRequest))
	}
}

func Round(request *Request) Response {
	url := random.AlphaNumeric(6)

	roundToPersist := round.UpdatedRoundWithNextCandidate(&persistence.Round{
		Url:          url,
		CreateDate:   time.Now(),
		UpdateDate:   time.Now(),
		Participants: transformParticipants(request),
	})

	persistence.CreateRound(roundToPersist)
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

func newRoundResponse(persistedRound *persistence.Round) Response {
	encodedRoundToken := encodeRoundToken(persistedRound)
	return Response{
		Token: &encodedRoundToken,
		Round: round.TransformRound(persistedRound),
	}
}

func encodeRoundToken(persistedRound *persistence.Round) string {
	return jwt.NewHelper().Encode(&jwt.RoundToken{
		RoundUrl: persistedRound.Url,
		StandardClaims: jwtLib.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	})
}
