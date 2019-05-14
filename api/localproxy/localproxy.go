package main

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/pointers"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

func main() {
	errors.LogFatal(app.WithHandlers(app.ApplicationHandlers{
		CreateRound: func(request *create.Request) create.Response {
			return create.Response{
				Token: pointers.String("daskdsa"),
				Round: &round.Round{
					Url: "theberesford.diet",
					Participants: []round.Participant{{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					}}}}
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
		JoinRound: func(context *gin.Context) {},
	}))
}
