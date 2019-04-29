package create

import (
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/random"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Participants []string `json:"participants"`
}

type Response struct {
	Url string `json:"url"`
}

func CreateRound(ctx *gin.Context) {
	createRoundRequest := new(Request)
	errors.LogFatal(ctx.ShouldBindJSON(createRoundRequest))
	url := random.RandomAlphaNumeric(6)
	persistence.CreateRound(url, createRoundRequest.Participants)
	ctx.JSON(200, Response{
		Url: url,
	})
}
