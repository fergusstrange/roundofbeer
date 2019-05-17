package app

import (
	"fmt"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
	"github.com/fergusstrange/roundofbeer/api/handlers/next"
	"github.com/fergusstrange/roundofbeer/api/handlers/read"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
	"os"
)

type ApplicationHandlers struct {
	CreateRound        func(*create.Request) round.WithToken
	JoinRound          func(roundId string, request *join.RoundRequest) (*round.WithToken, int)
	GetRound           func(roundToken string) (*round.Round, int)
	NextRoundCandidate func(roundToken string) (*round.Round, int)
}

func DefaultApp() error {
	return WithHandlers(DefaultHandlers())
}

func WithHandlers(handlers ApplicationHandlers) error {
	app := gin.Default()

	app.POST("/round", create.NewRoundHandler(handlers.CreateRound))
	app.POST("/round/:roundId", join.NewJoinRoundHandler(handlers.JoinRound))
	app.GET("/round", read.NewReadRoundHandler(handlers.GetRound))
	app.PUT("/round", next.NewNextRoundHandler(handlers.NextRoundCandidate))

	return app.Run(portFromEnvironment())
}

func DefaultHandlers() ApplicationHandlers {
	return ApplicationHandlers{
		CreateRound:        create.Round,
		JoinRound:          join.Round,
		GetRound:           read.Round,
		NextRoundCandidate: next.Round,
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
