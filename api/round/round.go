package round

import (
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"time"
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

func NewRoundResponse(round *persistence.Round) Response {
	return Response{
		Token: jwt.NewHelper().Encode(&jwt.RoundToken{
			RoundUrl: round.Url,
			StandardClaims: jwtLib.StandardClaims{
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			},
		}),
		Round: *TransformRound(round),
	}
}

func TransformRound(round *persistence.Round) *Round {
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
