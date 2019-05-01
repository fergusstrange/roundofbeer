package create

import (
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/random"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Participants []string `json:"participants"`
}

func NewRound(ctx *gin.Context) {
	createRoundRequest := new(Request)
	errors.LogError(ctx.BindJSON(createRoundRequest))
	url := random.AlphaNumeric(6)
	persistence.CreateRound(url, createRoundRequest.Participants)
	round := persistence.FetchRound(url)
	ctx.JSON(200, validation.NewRoundResponse(round))
}
