package create

import (
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/random"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Participants []string `json:"participants"`
}

func NewRoundHandler(serviceHandler func(request *Request) round.Response) func(*gin.Context) {
	return func(ctx *gin.Context) {
		createRoundRequest := new(Request)
		errors.LogError(ctx.BindJSON(createRoundRequest))
		ctx.JSON(200, serviceHandler(createRoundRequest))
	}
}

func NewRound(request *Request) round.Response {
	url := random.AlphaNumeric(6)
	persistence.CreateRound(url, request.Participants)
	fetchedRound := persistence.FetchRound(url)
	return round.NewRoundResponse(fetchedRound)
}
