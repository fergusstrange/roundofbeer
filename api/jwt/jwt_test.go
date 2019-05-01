package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelper_Encode(t *testing.T) {
	token := helperWithTestKey().Encode(&RoundToken{
		RoundUrl: "pdjasdnks",
	})

	assert.NotEmpty(t, token)
}

func TestHelper_Decode(t *testing.T) {
	token := helperWithTestKey().Encode(&RoundToken{
		RoundUrl: "pdjasdnks",
	})

	decodedRoundToken, err := helperWithTestKey().Decode(token)

	assert.NoError(t, err)
	assert.EqualValues(t,
		RoundToken{
			RoundUrl: "pdjasdnks",
		},
		decodedRoundToken,
	)
}

func helperWithTestKey() *Helper {
	helper := NewHelper()
	helper.SigningKey = []byte("abcdefg")
	return helper
}
