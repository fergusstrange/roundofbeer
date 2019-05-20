package main

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
	"github.com/fergusstrange/roundofbeer/api/pointers"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
)

func main() {
	errors.LogFatal(app.WithHandlers(app.ApplicationModule{
		Persistence: testfixtures.MockPersistence{},
		CreateRound: func(request *create.Request) round.WithToken {
			return round.WithToken{
				Token:    pointers.String("daskdsa"),
				RoundUrl: pointers.String("aUrl"),
			}
		},
		NextRoundCandidate: func(roundToken string) (*round.Round, int) {
			return &round.Round{
				Url: "theberesford.diet",
				Participants: []round.Participant{
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663dd",
						Name:       "Graeme",
						RoundCount: 103,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663aa",
						Name:       "Bert",
						RoundCount: 11,
					},
				}}, 200
		},
		GetRound: func(roundToken string) (*round.Round, int) {
			return &round.Round{
				Url: "theberesford.diet",
				Participants: []round.Participant{{
					UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
					Name:       "Tom",
					RoundCount: 33,
				}}}, 200
		},
		JoinRound: func(roundId string, request *join.RoundRequest) (*round.WithToken, int) {
			return &round.WithToken{
				Token:    pointers.String("daskdsa"),
				RoundUrl: pointers.String("theberesford.diet"),
			}, 200
		},
	}))
}
