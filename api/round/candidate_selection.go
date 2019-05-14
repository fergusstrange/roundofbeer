package round

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"math/rand"
	"time"
)

func UpdatedRoundWithNextCandidate(currentRound *persistence.Round) *persistence.Round {
	chosenParticipant := selectNextCandidate(currentRound)
	chosenParticipant.RoundCount = chosenParticipant.RoundCount + 1
	updatedParticipants := remapParticipantsWithChosen(currentRound, chosenParticipant)
	currentRound.CurrentCandidate = chosenParticipant.UUID
	currentRound.Participants = updatedParticipants
	return currentRound
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
	return candidatesForNextRound[randomIndexOrFirstWhenOnlyOneCandidate(candidatesForNextRound)]
}

func remapParticipantsWithChosen(round *persistence.Round, chosenParticipant persistence.Participant) []persistence.Participant {
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
