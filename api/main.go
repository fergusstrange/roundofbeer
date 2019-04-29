package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/fergusstrange/roundofbeer/api/handlers"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	log.SetHandler(text.Default)
	go persistence.CreateRoundTable()
}

func main() {
	app := gin.Default()

	app.POST("/round", create.CreateRound)
	app.GET("/round", handlers.GetRound)
	app.PUT("/round", handlers.IncrementRoundParticipant)

	if err := http.ListenAndServe(":3000", app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
