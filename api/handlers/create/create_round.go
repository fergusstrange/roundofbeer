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

func NewRoundHandler(serviceHandler func(request *Request) validation.Response) func(*gin.Context) {
	return func(ctx *gin.Context) {
		createRoundRequest := new(Request)
		errors.LogError(ctx.BindJSON(createRoundRequest))
		ctx.JSON(200, serviceHandler(createRoundRequest))
	}
}

func NewRound(request *Request) validation.Response {
	url := random.AlphaNumeric(6)
	persistence.CreateRound(url, request.Participants)
	round := persistence.FetchRound(url)
	return validation.NewRoundResponse(round)
}
