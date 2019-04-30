package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/read"
	"github.com/fergusstrange/roundofbeer/api/handlers/update"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetHandler(text.Default)
	persistence.CreateRoundTable()
}

func main() {
	app := gin.Default()

	app.POST("/round", create.CreateRound)
	app.GET("/round/:roundId", read.GetRound)
	app.PUT("/round/:roundId", update.IncrementRoundParticipant)

	errors.LogFatal(app.Run())
}
