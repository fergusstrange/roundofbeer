package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
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

	app.POST("/round", create.NewRound)
	app.POST("/round/:roundId", join.Round)
	app.GET("/round", read.Round)
	app.PUT("/round", update.Round)

	errors.LogFatal(app.Run())
}
