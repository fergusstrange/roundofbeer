package validation

import (
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Round struct {
	Url          string        `json:"url"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	RoundCount int    `json:"round_count"`
}

type Response struct {
	Token string `json:"token"`
	Round Round  `json:"round"`
}

func ValidRoundPathParam(ctx *gin.Context) string {
	roundId := ctx.Param("roundId")
	if roundId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundId
}

func ValidRoundHeader(ctx *gin.Context) string {
	roundToken := ctx.GetHeader("x-round-token")
	if roundToken == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	return roundToken
}

func NewRoundResponse(round *persistence.Round) Response {
	return Response{
		Token: jwt.NewHelper().Encode(jwt.RoundToken{
			RoundUrl: round.Url,
		}),
		Round: *Transform(round),
	}
}

func Transform(round *persistence.Round) *Round {
	var participants []Participant
	for _, participant := range round.Participants {
		participants = append(participants, Participant{
			UUID:       participant.UUID,
			Name:       participant.Name,
			RoundCount: participant.RoundCount,
		})
	}

	return &Round{
		Url:          round.Url,
		Participants: participants,
	}
}
