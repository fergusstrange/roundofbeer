package app

import (
	"fmt"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/join"
	"github.com/fergusstrange/roundofbeer/api/pointers"
	"github.com/fergusstrange/roundofbeer/api/round"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestPortFromEnvironment_Override(t *testing.T) {
	defer func() { assert.Nil(t, os.Unsetenv("PORT")) }()

	assert.Nil(t, os.Setenv("PORT", "3000"))

	actualPort := portFromEnvironment()

	assert.Equal(t, "localhost:3000", actualPort)
}

func TestPortFromEnvironment_Default(t *testing.T) {
	actualPort := portFromEnvironment()

	assert.Equal(t, "localhost:8080", actualPort)
}

func Test_VerifyProviderTests(t *testing.T) {
	go func() {
		errors.LogFatal(NewApplication(MockHandlers(), testfixtures.NewMockPersistence()).Run())
	}()

	provider := &dsl.Pact{
		Provider:                 "api",
		DisableToolValidityCheck: true,
	}

	dir, err := os.Getwd()
	assert.Nil(t, err)

	fmt.Println(dir)

	abs, err := filepath.Abs(dir)
	assert.Nil(t, err)

	fmt.Println(abs)

	pactFile := filepath.Join(filepath.Dir(abs), "ui/pacts/ui-api.json")
	fmt.Println(pactFile)

	pactFile, err = filepath.Abs("../../ui/pacts/ui-api.json")
	assert.Nil(t, err)

	_, err = provider.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		PactURLs:        []string{pactFile},
	})
	assert.Nil(t, err)
}

func MockHandlers() *ApplicationModule {
	return &ApplicationModule{
		CreateRound: func(request *create.Request) round.WithToken {
			return round.WithToken{
				Token:        pointers.String("daskdsa"),
				RoundURL:     pointers.String("aUrl"),
				Participants: []string{"Greg", "James"},
			}
		},
		NextRoundCandidate: func(roundToken string) (*round.Round, int) {
			return &round.Round{
				URL: "theberesford.diet",
				Participants: []round.Participant{
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663dd",
						Name:       "Graeme",
						RoundCount: 103,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663aa",
						Name:       "Bert",
						RoundCount: 11,
					},
				},
				CurrentCandidate: round.Participant{
					UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
					Name:       "Tom",
					RoundCount: 33,
				}}, 200
		},
		GetRound: func(roundToken string) (*round.Round, int) {
			return &round.Round{
				URL: "theberesford.diet",
				Participants: []round.Participant{
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663dd",
						Name:       "Graeme",
						RoundCount: 103,
					},
					{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663aa",
						Name:       "Bert",
						RoundCount: 11,
					},
				},
				CurrentCandidate: round.Participant{
					UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
					Name:       "Tom",
					RoundCount: 33,
				}}, 200
		},
		JoinRound: func(roundID string, request *join.RoundRequest) (*round.WithToken, int) {
			return &round.WithToken{
				Token:        pointers.String("daskdsa"),
				RoundURL:     pointers.String("aUrl"),
				Participants: []string{"Greg", "James"},
			}, 200
		},
	}
}
