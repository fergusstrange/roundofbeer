package update

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type Response struct {
	Participant string `json:"participant"`
}

func IncrementRoundParticipant(ctx *gin.Context) {
	roundId := validation.ValidRoundId(ctx)
	round := persistence.FetchRound(roundId)

	firstParticipant := round.Participants[0]
	lowestBought := firstParticipant.RoundCount
	var candidatesForNextRound = []persistence.Participant{firstParticipant}
	for _, participant := range round.Participants[1:] {
		if participant.RoundCount < lowestBought {
			lowestBought = participant.RoundCount
			candidatesForNextRound = []persistence.Participant{participant}
		} else if participant.RoundCount == lowestBought {
			candidatesForNextRound = append(candidatesForNextRound, participant)
		}
	}
	chosenParticipant := candidatesForNextRound[rand.Intn(len(candidatesForNextRound)-1)]
	chosenParticipant.RoundCount = chosenParticipant.RoundCount + 1

	var updatedParticipants []persistence.Participant
	for _, existingParticipant := range round.Participants {
		if existingParticipant.UUID == chosenParticipant.UUID {
			updatedParticipants = append(updatedParticipants, chosenParticipant)
		} else {
			updatedParticipants = append(updatedParticipants, existingParticipant)
		}
	}
	round.Participants = updatedParticipants
	persistence.UpdateRound(round)

	ctx.JSON(200, &Response{
		Participant: chosenParticipant.Name,
	})
}
