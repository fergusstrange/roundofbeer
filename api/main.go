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
		db := dynamo.New(newSession)
		err := db.CreateTable("roundofbeer", Round{}).Run()
		if err != nil {
			log.WithError(err).Fatalf("Unable to create DynamoDB table")
		}
	}
}

func main() {
	addr := ":" + os.Getenv("PORT")
	app := pat.New()
	app.Get("/round", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("{}"))
		if err != nil {
			log.Fatalf("Cannot write bytes")
		}
		writer.WriteHeader(200)
	})
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("error listening")
	}
}
