package read

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
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

func GetRound(ctx *gin.Context) {
	roundId := validation.ValidRoundId(ctx)
	round := persistence.FetchRound(roundId)
	ctx.JSON(200, transform(round))
}

func transform(round persistence.Round) *Round {
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
