package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/guregu/dynamo"
)

var Client = fetchClient()

func fetchClient() *dynamo.DB {
	newSession, err := session.NewSession()
	errors.LogFatal(err)
	return dynamo.New(newSession)
}
