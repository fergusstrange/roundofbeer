package update

import (
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type Response struct {
	Participant string `json:"participant"`
}

func Round(ctx *gin.Context) {
	roundToken, err := jwt.NewHelper().Decode(validation.ValidRoundHeader(ctx))
	if err != nil {
		ctx.AbortWithStatus(400)
	} else if round := persistence.FetchRound(roundToken.RoundUrl); round != nil {
		chosenParticipant := selectNextCandidate(round)
		chosenParticipant.RoundCount = chosenParticipant.RoundCount + 1
		updatedParticipants := updateParticipantsWithChosen(round, chosenParticipant)
		persistence.UpdateParticipants(round.Url, updatedParticipants)
		ctx.JSON(200, &Response{
			Participant: chosenParticipant.Name,
		})
	} else {
		ctx.AbortWithStatus(400)
	}
}

func selectNextCandidate(round *persistence.Round) persistence.Participant {
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
	chosenParticipant := candidatesForNextRound[randomIndexOrFirstWhenOnlyOneCandidate(candidatesForNextRound)]
	return chosenParticipant
}

func updateParticipantsWithChosen(round *persistence.Round, chosenParticipant persistence.Participant) []persistence.Participant {
	var updatedParticipants []persistence.Participant
	for _, existingParticipant := range round.Participants {
		if existingParticipant.UUID == chosenParticipant.UUID {
			updatedParticipants = append(updatedParticipants, chosenParticipant)
		} else {
			updatedParticipants = append(updatedParticipants, existingParticipant)
		}
	}
	return updatedParticipants
}

func randomIndexOrFirstWhenOnlyOneCandidate(candidatesForNextRound []persistence.Participant) int {
	numberOfCandidates := len(candidatesForNextRound)
	if numberOfCandidates <= 1 {
		return 0
	}
	return rand.New(rand.NewSource(time.Now().Unix())).Intn(numberOfCandidates)
}
