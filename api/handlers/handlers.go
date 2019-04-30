package handlers

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
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

func Transform(round persistence.Round) Round {
	var participants []Participant
	for _, participant := range round.Participants {
		participants = append(participants, Participant{
			UUID:       participant.UUID,
			Name:       participant.Name,
			RoundCount: participant.RoundCount,
		})
	}

	return Round{
		Url:          round.Url,
		Participants: participants,
	}
}
