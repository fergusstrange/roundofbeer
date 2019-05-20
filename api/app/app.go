package app

import (
	"fmt"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
	"github.com/fergusstrange/roundofbeer/api/handlers/next"
	"github.com/fergusstrange/roundofbeer/api/handlers/read"
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
	"os"
)

type ApplicationModule struct {
	Persistence        persistence.Persistence
	CreateRound        func(request *create.Request) round.WithToken
	JoinRound          func(roundId string, request *join.RoundRequest) (*round.WithToken, int)
	GetRound           func(roundToken string) (*round.Round, int)
	NextRoundCandidate func(roundTokenHeader string) (*round.Round, int)
}

func DefaultApp() error {
	return WithHandlers(DefaultModule())
}

func WithHandlers(applicationModule ApplicationModule) error {
	app := gin.Default()

	app.POST("/round", create.Handler(applicationModule.CreateRound))
	app.POST("/round/:roundId", join.Handler(applicationModule.JoinRound))
	app.GET("/round", read.Handler(applicationModule.GetRound))
	app.PUT("/round", next.Handler(applicationModule.NextRoundCandidate))

	applicationModule.Persistence.CreateRoundTable()

	return app.Run(portFromEnvironment())
}

func DefaultModule() ApplicationModule {
	dynamoDBPersistence := persistence.NewDynamoDBPersistence()
	return ApplicationModule{
		Persistence:        dynamoDBPersistence,
		CreateRound:        create.NewServiceContext(dynamoDBPersistence).ServiceHandler,
		JoinRound:          join.NewServiceContext(dynamoDBPersistence).ServiceHandler,
		GetRound:           read.NewServiceContext(dynamoDBPersistence).ServiceHandler,
		NextRoundCandidate: next.NewServiceContext(dynamoDBPersistence).ServiceHandler,
	}
}

func portFromEnvironment() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "localhost:8080"
	} else {
		return fmt.Sprintf(":%s", port)
	}
}
