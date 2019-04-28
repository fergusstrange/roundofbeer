package main

import (
	"github.com/apex/log/handlers/text"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/pat"
)

type Round struct {
	Id string `dynamo:"ID,hash"`
}

func init() {
	log.SetHandler(text.Default)
	newSession, err := session.NewSession()
	if err != nil {
		log.WithError(err).Fatalf("Unable to create DynamoDB table")
	} else {
		dynamo.New(newSession).CreateTable("roundofbeer", Round{})
	}
}

// setup.
func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
