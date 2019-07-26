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

type Application struct {
	persistence       persistence.Persistence
	applicationModule *ApplicationModule
	handlers          *gin.Engine
}

func NewApplication(applicationModule *ApplicationModule, persistence persistence.Persistence) *Application {
	return &Application{
		persistence:       persistence,
		applicationModule: applicationModule,
		handlers:          Handlers(applicationModule),
	}
}

func (application *Application) Initialise() {
	application.persistence.CreateRoundTable()
}

func (application *Application) Run() error {
	return application.handlers.Run(portFromEnvironment())
}

type ApplicationModule struct {
	CreateRound        func(request *create.Request) round.WithToken
	JoinRound          func(roundID string, request *join.RoundRequest) (*round.WithToken, int)
	GetRound           func(roundToken string) (*round.Round, int)
	NextRoundCandidate func(roundTokenHeader string) (*round.Round, int)
}

func NewApplicationModule(persistence persistence.Persistence) *ApplicationModule {
	return &ApplicationModule{
		CreateRound:        create.NewServiceContext(persistence).ServiceHandler,
		JoinRound:          join.NewServiceContext(persistence).ServiceHandler,
		GetRound:           read.NewServiceContext(persistence).ServiceHandler,
		NextRoundCandidate: next.NewServiceContext(persistence).ServiceHandler,
	}
}

func Handlers(applicationModule *ApplicationModule) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.POST("/round", create.Handler(applicationModule.CreateRound))
	app.POST("/round/:roundID", join.Handler(applicationModule.JoinRound))
	app.GET("/round", read.Handler(applicationModule.GetRound))
	app.PUT("/round", next.Handler(applicationModule.NextRoundCandidate))

	return app
}

func portFromEnvironment() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "localhost:8080"
	}
	return fmt.Sprintf("localhost:%s", port)
}
