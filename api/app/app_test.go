package app

import (
	"fmt"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/handlers/create"
	"github.com/fergusstrange/roundofbeer/api/handlers/validation"
	"github.com/gin-gonic/gin"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_VerifyProviderTests(t *testing.T) {
	go func() {
		errors.LogFatal(WithHandlers(MockHandlers()))
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

func MockHandlers() ApplicationHandlers {
	return ApplicationHandlers{
		CreateRound: func(request *create.Request) validation.Response {
			return validation.Response{
				Token: "daskdsa",
				Round: validation.Round{
					Url: "theberesford.diet",
					Participants: []validation.Participant{{
						UUID:       "d197f52e-5f9d-4082-92d7-fcbadf4663af",
						Name:       "Tom",
						RoundCount: 33,
					}}}}
		},
		NextRoundCandidate: func(context *gin.Context) {},
		GetRound:           func(context *gin.Context) {},
		JoinRound:          func(context *gin.Context) {},
	}
}
