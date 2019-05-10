package app

import (
	"fmt"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
	"github.com/fergusstrange/roundofbeer/api/handlers/read"
	"github.com/fergusstrange/roundofbeer/api/handlers/update"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/gin-gonic/gin"
	"os"
)

type ApplicationHandlers struct {
	CreateRound        func(*create.Request) round.Response
	JoinRound          func(*gin.Context)
	GetRound           func(*gin.Context)
	NextRoundCandidate func(*gin.Context)
}

func DefaultApp() error {
	return WithHandlers(DefaultHandlers())
}

func WithHandlers(handlers ApplicationHandlers) error {
	app := gin.Default()

	app.POST("/round", create.NewRoundHandler(handlers.CreateRound))
	app.POST("/round/:roundId", handlers.JoinRound)
	app.GET("/round", handlers.GetRound)
	app.PUT("/round", handlers.NextRoundCandidate)

	return app.Run(portFromEnvironment())
}

func DefaultHandlers() ApplicationHandlers {
	return ApplicationHandlers{
		CreateRound:        create.NewRound,
		JoinRound:          join.Round,
		GetRound:           read.Round,
		NextRoundCandidate: update.Round,
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
