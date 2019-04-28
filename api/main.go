package main

import (
	"github.com/apex/log/handlers/text"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fergusstrange/roundofbeer/api/handlers"
	"github.com/guregu/dynamo"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

type Round struct {
	Id string `dynamo:"ID,hash"`
}

func init() {
	log.SetHandler(text.Default)
	newSession, err := session.NewSession()
	if err != nil {
		log.WithError(err).Error("Unable to create DynamoDB table")
	} else {
		db := dynamo.New(newSession)
		err := db.CreateTable("roundofbeer", Round{}).
			Index(dynamo.Index{
			}).
			Provision(5, 5).
			Run()
		if err != nil {
			log.WithError(err).Error("Unable to create DynamoDB table")
		}
	}
}

func main() {
	addr := ":" + os.Getenv("PORT")
	app := gin.Default()

	app.POST("/round", handlers.CreateRound)
	app.GET("/round", handlers.GetRound)
	app.PUT("/round", handlers.IncrementRoundParticipant)
	app.GET("/rounds", handlers.GetRoundsByParticipant)

	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
