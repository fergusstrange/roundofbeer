package handlers

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	"testing"
)

func TestCreateRound(t *testing.T) {
	handlers := app.Handlers(app.NewApplicationModule(testfixtures.NewMockPersistence()))

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(handlers).
		Post("/round").
		Body(`{"participants":["John","James","Steve"]}`).
		Expect(t).
		Assert(jsonpath.Matches("$.token", `^[\w\d]{36}\.[\w\d]{74}\.[\w\d-]{86}$`)).
		Assert(jsonpath.Matches("$.roundUrl", `^[\w\d]{6}$`)).
		Assert(jsonpath.Len("$.participants", 3)).
		Status(http.StatusOK).
		End()
}

func TestJoinRound(t *testing.T) {
	persistence := testfixtures.NewMockPersistence()
	handlers := app.Handlers(app.NewApplicationModule(persistence))

	persistence.NewTestRound("12345")

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(handlers).
		Post("/round/12345").
		Body(`{"name":"tom"}`).
		Expect(t).
		Assert(jsonpath.Matches("$.token", `^[\w\d]{36}\.[\w\d]{72}\.[\w\d-]{86}$`)).
		Status(http.StatusOK).
		End()
}

func TestNextRound(t *testing.T) {
	module := app.NewApplicationModule(testfixtures.NewMockPersistence())
	handlers := app.Handlers(module)

	roundToken := module.CreateRound(&create.Request{
		Participants: []string{"Tom", "Bert"},
	})

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(handlers).
		Put("/round").
		Header("x-round-token", *roundToken.Token).
		Expect(t).
		Assert(jsonpath.Matches("$.url", `^[\w\d]{6}$`)).
		Assert(jsonpath.Len("$.participants", 2)).
		Assert(jsonpath.Present("$.currentCandidate")).
		Status(http.StatusOK).
		End()
}

func TestFetchRound(t *testing.T) {
	module := app.NewApplicationModule(testfixtures.NewMockPersistence())
	handlers := app.Handlers(module)

	roundToken := module.CreateRound(&create.Request{
		Participants: []string{"Tom", "Bert"},
	})

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(handlers).
		Get("/round").
		Header("x-round-token", *roundToken.Token).
		Expect(t).
		Assert(jsonpath.Matches("$.url", `^[\w\d]{6}$`)).
		Assert(jsonpath.Len("$.participants", 2)).
		Assert(jsonpath.Present("$.currentCandidate")).
		Status(http.StatusOK).
		End()
}
