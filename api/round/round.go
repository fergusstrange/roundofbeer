package round

import (
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"time"
)

type WithToken struct {
	Token *string `json:"token"`
	RoundURL *string `json:"roundUrl"`
	Participants []string `json:"participants"`
}

type Round struct {
	URL              string        `json:"url"`
	Participants     []Participant `json:"participants"`
	CurrentCandidate Participant   `json:"currentCandidate"`
}

type Participant struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	RoundCount int    `json:"roundCount"`
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
		URL:              round.URL,
		Participants:     participants,
		CurrentCandidate: currentCandidate,
	}
}

func EncodeRoundToken(roundURL string) string {
	return jwt.NewHelper().Encode(&jwt.RoundToken{
		RoundURL: roundURL,
		StandardClaims: jwtLib.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	})
}

func transformParticipant(participant persistence.Participant) Participant {
	return Participant{
		UUID:       participant.UUID,
		Name:       participant.Name,
		RoundCount: participant.RoundCount,
	}
}

func ParticipantsFromRound(round *persistence.Round) []string {
	participants := make([]string, 0)
	for _, participant := range round.Participants {
		participants = append(participants, participant.Name)
	}
	return participants
}