package round

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
)

type Round struct {
	Url              string        `json:"url"`
	Participants     []Participant `json:"participants"`
	CurrentCandidate Participant   `json:"current_candidate"`
}

type Participant struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	RoundCount int    `json:"round_count"`
}

func TransformRound(round *persistence.Round) *Round {
	var participants []Participant
	var currentCandidate Participant
	for _, participant := range round.Participants {
		transformedParticipant := transformParticipant(participant)
		participants = append(participants, transformedParticipant)
		if transformedParticipant.UUID == round.CurrentCandidate {
			currentCandidate = transformedParticipant
		}
	}

	return &Round{
		Url:              round.Url,
		Participants:     participants,
		CurrentCandidate: currentCandidate,
	}
}

func transformParticipant(participant persistence.Participant) Participant {
	return Participant{
		UUID:       participant.UUID,
		Name:       participant.Name,
		RoundCount: participant.RoundCount,
	}
}
