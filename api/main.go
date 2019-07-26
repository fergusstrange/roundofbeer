package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	gin.DefaultErrorWriter = os.Stderr
	gin.DefaultWriter = os.Stderr
	log.SetHandler(text.Default)
}

func main() {
	dynamoDBPersistence := persistence.NewDynamoDBPersistence()
	applicationModule := app.NewApplicationModule(dynamoDBPersistence)
	application := app.NewApplication(applicationModule, dynamoDBPersistence)
	errors.LogFatal(application.Run())
}
