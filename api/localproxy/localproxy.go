package main

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

func main() {
	errors.LogFatal(app.WithHandlers(app.ApplicationHandlers{
		CreateRound: func(request *create.Request) round.Response {
			return round.Response{
				Token: "daskdsa",
				Round: round.Round{
					Url: "theberesford.diet",
					Participants: []round.Participant{{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					}}}}
		},
		NextRoundCandidate: func(context *gin.Context) {},
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
